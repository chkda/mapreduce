package master

import (
	"context"
	"log"
	"sync"

	pbm "github.com/chkda/mapreduce/rpc/master"
	pbw "github.com/chkda/mapreduce/rpc/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	Workers        map[string]*Worker
	DFS            *DFS
	Mu             sync.Mutex
	MapTasks       map[string]*MapTask
	ActiveMapTasks int
}

type Worker struct {
	Uuid   string
	IP     string
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
		Uuid: workerInfo.GetUuid(),
		IP:   workerInfo.GetIp(),
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
	mapTask := s.MapTasks[mapTaskId]
	mapTask.TaskStatus = COMPLETE
	mapTask.Result = mapResult.GetFilenames()
	s.ActiveMapTasks--
	s.Mu.Unlock()
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
	go s.processData(filename)
	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) processData(filename string) {
	s.startMapPhase(filename)
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
			s.ActiveMapTasks++
			go func(task *MapTask, client pbw.WorkerClient) {
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
				return
			}(mapTask, workerClient)
		}
	}
}
