package master

import (
	"context"
	"errors"
	"hash/fnv"
	"log"
	"sync"
	"time"

	pbm "github.com/chkda/mapreduce/rpc/master"
	pbw "github.com/chkda/mapreduce/rpc/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerStatus int

const (
	WORKER_ALIVE WorkerStatus = iota
	WORKER_DEAD
)

var (
	ErrTaskIdDoesnotExist         = errors.New("taskId doesn't exist")
	ErrDeadWorkerThresholdReached = errors.New("dead workers have exceeded more than threshold")
)

type Service struct {
	Workers              map[string]*Worker
	DFS                  *DFS
	Mu                   sync.RWMutex
	MapTasks             map[string]*MapTask
	DeadWorkersThreshold int
	DeadWorkers          []string
	NumReduce            int
	ActiveMapTasks       int
}

type Worker struct {
	Uuid   string
	IP     string
	State  WorkerStatus
	Client pbw.WorkerClient
}

func New() *Service {
	dfs := NewDFS()
	return &Service{
		Workers: make(map[string]*Worker),
		DFS:     dfs,
	}
}

func (s *Service) RegisterWorker(ctx context.Context, workerInfo *pbm.WorkerInfo) (*pbm.Ack, error) {
	worker := &Worker{
		Uuid:  workerInfo.GetUuid(),
		IP:    workerInfo.GetIp(),
		State: WORKER_ALIVE,
	}
	conn, err := grpc.Dial(worker.IP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbw.NewWorkerClient(conn)
	worker.Client = client
	s.Mu.Lock()
	s.Workers[worker.Uuid] = worker
	s.Mu.Unlock()
	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) UpdateMapResult(ctx context.Context, mapResult *pbm.MapResult) (*pbm.Ack, error) {
	mapTaskId := mapResult.GetUuid()
	s.Mu.Lock()
	defer s.Mu.Unlock()
	mapTask, ok := s.MapTasks[mapTaskId]
	if !ok {
		return nil, ErrTaskIdDoesnotExist
	}
	mapTask.TaskStatus = COMPLETE
	mapTask.OutputFiles = mapResult.GetFilenames()
	s.ActiveMapTasks--
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
	filename := taskRequest.GetFilename()
	s.Mu.Lock()
	s.NumReduce = int(taskRequest.GetNumReduce())
	s.Mu.Unlock()
	go s.processData(filename)
	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) processData(filename string) {
	s.startMapPhase(filename)

	for {
		// TODO : check active worker, map task states and reduce states
		s.checkActiveWorkers()
		s.Mu.Lock()
		deadWorkers := s.DeadWorkers
		s.Mu.Unlock()
		if len(deadWorkers) > s.DeadWorkersThreshold {
			log.Println("Dead Workers:", deadWorkers)
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

func (s *Service) startMapPhase(filename string) {
	dataNodes, err := s.DFS.GetDataNodes(filename)
	if err != nil {
		log.Println(err)
		return
	}
	// failureChan := make(chan string)

	for _, node := range dataNodes {
		workerClient := s.Workers[node.Uuid].Client
		for i := 0; i < len(node.Filenames); i++ {
			mapTask := NewMapTask(node.Filenames[i], node.Uuid)
			s.MapTasks[mapTask.ID] = mapTask
			go s.startMapTask(mapTask, workerClient)
		}
	}
}

func (s *Service) startMapTask(task *MapTask, client pbw.WorkerClient) {
	mapRequest := &pbw.MapTask{
		TaskId:    task.ID,
		Filename:  task.TaskFile,
		NumReduce: 0, // TODO : Look at this later
	}
	ack, err := client.AssignMap(context.Background(), mapRequest)
	if err != nil || !ack.Success {
		task.TaskStatus = FAILED
		return
	}
	task.TaskStatus = INPROGRESS
	s.Mu.Lock()
	s.ActiveMapTasks++
	s.Mu.Unlock()
	return
}

func (s *Service) startReducePhase() {
	mapTasksCopy := make(map[string]*MapTask)
	s.Mu.RLock()
	for k, v := range s.MapTasks {
		mapTasksCopy[k] = v
	}
	s.Mu.RUnlock()

	// workerIDHashMap := s.calculateWorkerIndex()
}

func (s *Service) calculateWorkerIndex() map[uint32]string {
	workerIDHashMap := make(map[uint32]string)
	s.Mu.RLock()
	for workerId, _ := range s.Workers {
		hashId := hash(workerId) % uint32(s.NumReduce)
		workerIDHashMap[hashId] = workerId
	}
	s.Mu.RUnlock()
	return workerIDHashMap
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
	client := worker.Client
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
	s.Mu.Lock()
	defer s.Mu.Unlock()

	worker, ok := s.Workers[workerId]
	if !ok {
		return
	}
	if alive {
		worker.State = WORKER_ALIVE
		return
	}
	worker.State = WORKER_DEAD
	s.DeadWorkers = append(s.DeadWorkers, workerId)
	return
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
