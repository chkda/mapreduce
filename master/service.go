package master

import (
	"context"
	"errors"
	"hash/fnv"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	pbm "github.com/chkda/mapreduce/rpc/master"
	pbw "github.com/chkda/mapreduce/rpc/worker"
)

var (
	ErrReducerIdNotInt            = errors.New("reducer id not an integer")
	ErrTaskIdDoesnotExist         = errors.New("taskId doesn't exist")
	ErrWorkerIdDoesnotExist       = errors.New("workerId doesn't exist")
	ErrWrongFilenameFormat        = errors.New("wrong filename format")
	ErrDeadWorkerThresholdReached = errors.New("dead workers have exceeded more than threshold")
)

type Service struct {
	Filename             string
	Workers              map[string]*Worker
	DFS                  *DFS
	Mu                   sync.RWMutex
	MapTasks             map[string]*MapTask
	ReduceTasks          map[string]*ReduceTask
	DeadWorkersThreshold int
	DeadWorkers          map[string]bool
	NumReduce            int
	ActiveMapTasks       int
	ActiveReduceTasks    int
}

func New(filename string, numReduce int, deadWorkerThreshold int) *Service {
	dfs := NewDFS()
	return &Service{
		Workers:              make(map[string]*Worker),
		DFS:                  dfs,
		Filename:             filename,
		NumReduce:            numReduce,
		DeadWorkersThreshold: deadWorkerThreshold,
	}
}

func (s *Service) GetFilename() string {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	return s.Filename
}

func (s *Service) AddWorker(id string, worker *Worker) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Workers[id] = worker
}

func (s *Service) GetWorker(workerId string) (*Worker, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	worker, ok := s.Workers[workerId]
	if !ok {
		return nil, ErrWorkerIdDoesnotExist
	}
	return worker, nil
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

func (s *Service) GetReduceTask(taskId string) (*ReduceTask, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	task, ok := s.ReduceTasks[taskId]
	if !ok {
		return nil, ErrTaskIdDoesnotExist
	}
	return task, nil
}

func (s *Service) IncActiveMapTasks() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.ActiveMapTasks++
}

func (s *Service) DecActiveMapTasks() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.ActiveMapTasks--
}

func (s *Service) IncActiveReduceTasks() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.ActiveReduceTasks++
}

func (s *Service) DecActiveReduceTasks() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.ActiveReduceTasks--
}

func (s *Service) AddDeadWorkers(workerId string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.DeadWorkers[workerId] = true
}

func (s *Service) RemoveDeadWorkers(workerId string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.DeadWorkers, workerId)
}

func (s *Service) checkIfDeadWorkersExceedsThreshold() bool {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	deadWorkers := s.DeadWorkers
	if len(deadWorkers) >= s.DeadWorkersThreshold {
		return true
	}
	return false
}

func (s *Service) RegisterWorker(ctx context.Context, workerInfo *pbm.WorkerInfo) (*pbm.Ack, error) {
	worker, err := NewWorker(
		workerInfo.GetUuid(),
		workerInfo.GetIp(),
	)

	if err != nil {
		return nil, err
	}

	s.AddWorker(workerInfo.GetUuid(), worker)
	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) UpdateMapResult(ctx context.Context, mapResult *pbm.MapResult) (*pbm.Ack, error) {
	mapTaskId := mapResult.GetUuid()
	mapTask, err := s.GetMapTask(mapTaskId)
	if err != nil {
		return nil, err
	}
	mapTask.SetTaskStatus(COMPLETE)
	mapTask.SetOutputFiles(mapResult.GetOutputFiles())
	s.IncActiveMapTasks()
	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) UpdateDataNodes(ctx context.Context, nodesInfo *pbm.DataNodesInfo) (*pbm.Ack, error) {
	nodes := nodesInfo.GetNodes()
	fileName := nodesInfo.GetFilename()
	datanodes := make([]*DataNode, 0, 4)
	for _, node := range nodes {
		dataNode := &DataNode{
			Uuid:      node.GetUuid(),
			Filenames: node.GetFiles(),
		}
		datanodes = append(datanodes, dataNode)
	}
	s.Mu.Lock()
	s.DFS.FileChunks[fileName] = datanodes
	s.Mu.Unlock()
	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) Trigger(ctx context.Context, taskRequest *pbm.TaskRequest) (*pbm.Ack, error) {
	go s.processData()
	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) processData() {
	s.startMapPhase()

	for {
		// TODO : check active worker, map task states and reduce states
		s.checkActiveWorkers()

		if s.checkIfDeadWorkersExceedsThreshold() {
			// log.Println("Dead Workers:", deadWorkers)
			log.Fatal(ErrDeadWorkerThresholdReached)
		}

		s.Mu.Lock()
		for _, task := range s.MapTasks {
			workerId := task.WorkerID
			worker := s.Workers[workerId]
			if worker.State == WORKER_DEAD {
				task.TaskStatus = FAILED
				continue
			}
		}
		s.Mu.Unlock()

	}
}

func (s *Service) startMapPhase() {
	filename := s.GetFilename()

	dataNodes, err := s.DFS.GetDataNodes(filename)
	if err != nil {
		log.Println(err)
		return
	}

	// failureChan := make(chan string)
	mapTasks := make(map[string]*MapTask)
	for _, node := range dataNodes {
		worker, err := s.GetWorker(node.Uuid)
		if err != nil {
			log.Println(err)
			log.Println("node id not present in worker map:", node.Uuid)
			continue
		}
		workerClient := worker.GetClient()
		for i := 0; i < len(node.Filenames); i++ {
			mapTask := NewMapTask(node.Filenames[i], node.Uuid, s.NumReduce)
			mapTasks[mapTask.ID] = mapTask
			go s.startMapTask(mapTask, workerClient)
		}
	}

	s.Mu.Lock()
	s.MapTasks = mapTasks
	s.Mu.Unlock()
}

func (s *Service) startMapTask(task *MapTask, client pbw.WorkerClient) {
	mapRequest := &pbw.MapTask{
		TaskId:    task.GetId(),
		Filename:  task.GetTaskFile(),
		NumReduce: int32(task.GetNumReduce()), // TODO : Look at this later
	}
	ack, err := client.AssignMap(context.Background(), mapRequest)
	if err != nil || !ack.Success {
		task.TaskStatus = FAILED
		return
	}
	task.SetTaskStatus(INPROGRESS)
	s.IncActiveMapTasks()
}

func (s *Service) startReducePhase() {
	reducerToWorkerIdMap := s.getReducerIdToWorker()
	reducerIdToFilePartitionsMap := s.getReducerIdToFilePations()
	reduceTasks := make(map[string]*ReduceTask)

	for reducerId, reducerFiles := range reducerIdToFilePartitionsMap {
		workerId, ok := reducerToWorkerIdMap[reducerId]
		if !ok {
			log.Println("reducerId not present in reducer to worker map:", reducerId)
			continue
		}
		worker, err := s.GetWorker(workerId)
		if err != nil {
			log.Println(err)
			log.Println("node id not present in worker map:", workerId)
			continue
		}
		workerClient := worker.GetClient()
		reduceTask := NewReduceTask(reducerFiles, workerId)
		reduceTasks[reduceTask.ID] = reduceTask
		go s.startReduceTask(reduceTask, workerClient)
	}

	s.Mu.Lock()
	s.ReduceTasks = reduceTasks
	s.Mu.Unlock()
}

func (s *Service) startReduceTask(task *ReduceTask, client pbw.WorkerClient) {
	dataNodes := make([]*pbw.NodeFileInfo, 0, 5)
	taskFiles := task.GetTaskFiles()
	for _, taskFile := range taskFiles {
		dataNodes = append(dataNodes, &pbw.NodeFileInfo{
			File: taskFile.Filename,
			Ip:   taskFile.NodeIP,
		})
	}
	reduceRequest := &pbw.ReduceTask{
		TaskId: task.ID,
		Datanodes: &pbw.DataNodesInfo{
			Nodes: dataNodes,
		},
	}
	ack, err := client.AssignReduce(context.Background(), reduceRequest)
	if err != nil || !ack.Success {
		task.TaskStatus = FAILED
		return
	}
	task.SetTaskStatus(INPROGRESS)
	s.IncActiveReduceTasks()
	return

}

func (s *Service) getReducerIdToWorker() map[uint32]string {
	reducerToWorkerIdMap := make(map[uint32]string)
	s.Mu.RLock()
	for workerId, _ := range s.Workers {
		hashId := hash(workerId) % uint32(s.NumReduce)
		reducerToWorkerIdMap[hashId] = workerId
	}
	s.Mu.RUnlock()
	return reducerToWorkerIdMap
}

func (s *Service) getReducerIdToFilePations() map[uint32][]*ReduceDataNodeInfo {
	reducerIdToFilePartitionsMap := make(map[uint32][]*ReduceDataNodeInfo)
	mapTasksCopy := make(map[string]*MapTask)
	s.Mu.RLock()
	for k, v := range s.MapTasks {
		mapTasksCopy[k] = v
	}
	s.Mu.RUnlock()
	for _, task := range mapTasksCopy {
		for _, filepartion := range task.OutputFiles {
			reducerId, err := s.calculateReducerIdFromFilePartition(filepartion)
			if err != nil {
				log.Printf("filename: %s : %s", filepartion, err.Error())
				continue
			}
			_, ok := reducerIdToFilePartitionsMap[reducerId]
			if !ok {
				reducerIdToFilePartitionsMap[reducerId] = make([]*ReduceDataNodeInfo, 0, 5)
			}
			files := reducerIdToFilePartitionsMap[reducerId]
			workerId := task.GetWorkerId()
			worker, err := s.GetWorker(workerId)
			if err != nil {
				log.Println(err)
				log.Println("node id not present in worker map:", task.WorkerID)
				continue
			}
			files = append(files, &ReduceDataNodeInfo{
				Filename: filepartion,
				NodeIP:   worker.IP,
			})
			reducerIdToFilePartitionsMap[reducerId] = files
		}
	}

	return reducerIdToFilePartitionsMap
}

func (s *Service) calculateReducerIdFromFilePartition(filename string) (uint32, error) {
	filename = strings.Trim(filename, filepath.Ext(filename))
	parts := strings.Split(filename, "-")
	if len(parts) != 3 || parts[0] != "mr" {
		return 0, ErrWrongFilenameFormat
	}

	reducerId, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, ErrReducerIdNotInt
	}
	return uint32(reducerId), nil
}

func (s *Service) checkActiveWorkers() {
	workerCopy := make(map[string]*Worker)
	s.Mu.RLock()
	for k, v := range s.Workers {
		workerCopy[k] = v
	}
	s.Mu.RUnlock()

	var wg sync.WaitGroup
	for workerId, worker := range workerCopy {
		wg.Add(1)
		go func(nodeId string, node *Worker) {
			defer wg.Done()
			isAlive := s.checkActiveWorker(node)
			s.updateWorkerState(nodeId, isAlive)
		}(workerId, worker)
	}
	wg.Wait()
}

func (s *Service) checkActiveWorker(worker *Worker) bool {
	client := worker.GetClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.HealthCheck(ctx, &pbw.HealthcheckRequest{})
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *Service) updateWorkerState(workerId string, alive bool) {
	worker, err := s.GetWorker(workerId)
	if err != nil {
		log.Println(err)
		return
	}
	if alive {
		worker.SetState(WORKER_ALIVE)
		return
	}
	worker.SetState(WORKER_DEAD)
	s.AddDeadWorkers(workerId)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
