package worker

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	pbw "github.com/chkda/mapreduce/rpc/worker"
)

var (
	ErrUnableToOpenFile   = errors.New("unable to open file")
	ErrUnableToReadFile   = errors.New("unable to read file")
	ErrTaskIdDoesnotExist = errors.New("taskId doesn't exist")
)

type Service struct {
	pbw.WorkerServer
	Mu          sync.RWMutex
	ID          string
	IP          string
	MasterIP    string
	MapTasks    map[string]*MapTask
	ReduceTasks map[string]*ReduceTask
}

func New(cfg *Config) *Service {
	return &Service{
		ID:       cfg.ID,
		IP:       cfg.IP,
		MasterIP: cfg.MasterIP,
	}
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
}

func (s *Service) executeReduceTask(taskId string) {
	// TODO
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
