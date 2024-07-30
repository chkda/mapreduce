package worker

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	pbm "github.com/chkda/mapreduce/rpc/master"
	pbw "github.com/chkda/mapreduce/rpc/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ErrUnableToOpenFile           = errors.New("unable to open file")
	ErrUnableToReadFile           = errors.New("unable to read file")
	ErrTaskIdDoesnotExist         = errors.New("taskId doesn't exist")
	ErrUnableToRegisterWithMaster = errors.New("unable to register with master")
)

type Service struct {
	pbw.WorkerServer
	Mu           sync.RWMutex
	ID           string
	IP           string
	MasterIP     string
	MasterClient pbm.MasterClient
	MasterConn   *grpc.ClientConn
	MapTasks     map[string]*MapTask
	ReduceTasks  map[string]*ReduceTask
}

func New(cfg *Config) (*Service, error) {

	worker := &Service{
		ID: cfg.ID,
		IP: cfg.IP,
		MasterIP: cfg.
			MasterIP,
	}

	conn, err := grpc.Dial(worker.MasterIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbm.NewMasterClient(conn)
	worker.MasterClient = client
	worker.MasterConn = conn
	request := &pbm.WorkerInfo{
		Uuid: cfg.ID,
		Ip:   cfg.IP,
	}
	ack, err := client.RegisterWorker(context.Background(), request)
	if err != nil || !ack.Success {
		return nil, ErrUnableToRegisterWithMaster
	}
	return worker, nil
}

func (s *Service) GetMasterClient() pbm.MasterClient {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	return s.MasterClient
}

func (s *Service) GetMapTask(taskId string) (*MapTask, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	task, ok := s.MapTasks[taskId]
	if !ok {
		return nil, ErrTaskIdDoesnotExist
	}
	return task, nil
}

func (s *Service) AddMapTask(taskId string, task *MapTask) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.MapTasks[taskId] = task
}

func (s *Service) GetReduceTask(taskId string) (*ReduceTask, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	task, ok := s.ReduceTasks[taskId]
	if !ok {
		return nil, ErrTaskIdDoesnotExist
	}
	return task, nil
}

func (s *Service) AddReduceTask(taskId string, task *ReduceTask) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.ReduceTasks[taskId] = task
}

func (s *Service) HealthCheck(ctx *context.Context, req *pbw.HealthcheckRequest) (*pbw.WorkerAck, error) {
	return &pbw.WorkerAck{
		Success: true,
	}, nil
}

func (s *Service) GetTaskStatus(ctx *context.Context, req *pbw.StatusRequest) (*pbw.StatusResponse, error) {
	if req.Type == pbw.TaskType_MAP {
		task, err := s.GetMapTask(req.GetTaskId())
		if err != nil {
			return nil, err
		}
		return &pbw.StatusResponse{
			TaskId: task.GetId(),
			Status: pbw.Status(task.GetTaskStatus()),
		}, nil
	}
	task, err := s.GetReduceTask(req.GetTaskId())
	if err != nil {
		return nil, err
	}
	return &pbw.StatusResponse{
		TaskId: task.GetId(),
		Status: pbw.Status(task.GetTaskStatus()),
	}, nil
}

func (s *Service) AssignMap(ctx context.Context, mapRequest *pbw.MapTask) (*pbw.WorkerAck, error) {
	task := NewMapTask(
		mapRequest.GetFilename(),
		mapRequest.GetTaskId(),
		int(mapRequest.GetNumReduce()),
	)
	s.AddMapTask(mapRequest.GetTaskId(), task)
	go s.executeMapTask(mapRequest.GetTaskId())
	return &pbw.WorkerAck{
		Success: true,
	}, nil
}

func (s *Service) AssignReduce(ctx context.Context, reduceRequest *pbw.ReduceTask) (*pbw.WorkerAck, error) {
	intermediateData := make([]*ReduceDataNodeInfo, 0, 5)
	datanodes := reduceRequest.GetDatanodes()
	for _, data := range datanodes.GetNodes() {
		intermediateData = append(intermediateData, &ReduceDataNodeInfo{
			Filename: data.GetFile(),
			NodeIP:   data.GetIp(),
		})
	}
	task := NewReduceTask(intermediateData, reduceRequest.GetTaskId())
	s.AddReduceTask(reduceRequest.GetTaskId(), task)
	go s.executeReduceTask(reduceRequest.GetTaskId())
	return &pbw.WorkerAck{
		Success: true,
	}, nil
}

func (s *Service) GetIntermediateData(ctx context.Context, req *pbw.InterMediateDataRequest) (*pbw.InterMediateDataResponse, error) {
	filename := req.GetFilename()
	content, err := s.readDataFromFile(filename)
	if err != nil {
		return nil, err
	}
	return &pbw.InterMediateDataResponse{
		Data: content,
	}, nil
}

func (s *Service) executeMapTask(taskId string) {
	// TODO
	task, err := s.GetMapTask(taskId)
	if err != nil {
		log.Println(err)
		return
	}

	taskFile := task.GetTaskFile()
	mapResult := &pbm.MapResult{
		Uuid: taskId,
	}
	contentBytes, err := s.readDataFromFile(taskFile)
	if err != nil {
		log.Println()
		mapResult.TaskStatus = pbm.Status_FAILED
		go s.updateMapResult(mapResult)
		return
	}
	numReduce := task.GetNumReduce()

	partitionFiles := make([]string, 0, numReduce)
	partitionWriters := make([]*bufio.Writer, 0, numReduce)

	for i := 0; i < numReduce; i++ {
		filename := fmt.Sprintf("mr-%s-%d.txt", taskId, i)
		partitionFiles[i] = filename
		file, err := os.Create(filename)
		if err != nil {
			log.Println(err)
			mapResult.TaskStatus = pbm.Status_FAILED
			go s.updateMapResult(mapResult)
			return
		}
		defer file.Close()
		partitionWriters[i] = bufio.NewWriter(file)
		defer partitionWriters[i].Flush()
	}

	s.Map(taskFile, string(contentBytes), func(key, value string) {
		partitionId := hash(key) % uint32(numReduce)
		writer := partitionWriters[partitionId]
		writer.WriteString(key)
		writer.WriteByte('\t')
		writer.WriteString(value)
		writer.WriteByte('\n')
	})

	for _, writer := range partitionWriters {
		err := writer.Flush()
		if err != nil {
			log.Println(err)
		}
	}
	task.SetOutputFiles(partitionFiles)
	task.SetTaskStatus(COMPLETED)
	mapResult.OutputFiles = partitionFiles
	go s.updateMapResult(mapResult)

}

func (s *Service) executeReduceTask(taskId string) {
	// TODO
	task, err := s.GetReduceTask(taskId)
	if err != nil {
		log.Println(err)
		return
	}

	taskFiles := task.GetTaskFiles()
	reduceResult := &pbm.ReduceResult{
		Uuid: taskId,
	}

	intermediateKV := make(map[string][]string)
	for _, dataInfo := range taskFiles {
		data, err := s.getIntermediateData(dataInfo)
		if err != nil {
			log.Println(err)
			reduceResult.TaskStatus = pbm.Status_FAILED
			go s.updateReduceResult(reduceResult)
			return
		}
		s.processIntermediateData(data, intermediateKV)
	}
}

func (s *Service) getIntermediateData(dataInfo *ReduceDataNodeInfo) ([]byte, error) {
	if dataInfo.NodeIP == s.IP {
		return s.readDataFromFile(dataInfo.Filename)
	}
	request := &pbw.InterMediateDataRequest{
		Filename: dataInfo.Filename,
	}

	conn, err := grpc.Dial(dataInfo.NodeIP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := pbw.NewWorkerClient(conn)
	response, err := client.GetIntermediateData(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (s *Service) processIntermediateData(data []byte, kvMap map[string][]string) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "\t")
		if len(parts) == 2 {
			key, value := parts[0], parts[1]
			if _, ok := kvMap[key]; !ok {
				kvMap[key] = make([]string, 0, 5)
			}
			kvMap[key] = append(kvMap[key], value)
		}
	}
}

func (s *Service) updateMapResult(result *pbm.MapResult) {
	maxRetry := 3
	client := s.GetMasterClient()
	ctx := context.Background()
	for i := 0; i < maxRetry; i++ {
		ack, err := client.UpdateMapResult(ctx, result)
		if err != nil || !ack.Success {
			log.Println(err)
			continue
		}
		break
	}
}

func (s *Service) updateReduceResult(result *pbm.ReduceResult) {
	maxRetry := 3
	client := s.GetMasterClient()
	ctx := context.Background()
	for i := 0; i < maxRetry; i++ {
		ack, err := client.UpdateReduceResult(ctx, result)
		if err != nil || !ack.Success {
			log.Println(err)
			continue
		}
		break
	}
}

func (s *Service) Map(key string, value string, emitIntermediatePair func(string, string)) {
	scanner := bufio.NewScanner(strings.NewReader(value))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		emitIntermediatePair(word, "1")
	}
}

func (s *Service) readDataFromFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrUnableToOpenFile
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, ErrUnableToReadFile
	}
	return content, nil
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
