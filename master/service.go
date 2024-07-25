package master

import (
	"context"
	"sync"

	pbm "github.com/chkda/mapreduce/rpc/master"
	pbw "github.com/chkda/mapreduce/rpc/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	Workers map[string]*Worker
	DFS     *DFS
	Mu      sync.Mutex
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
	defer s.Mu.Unlock()
	s.Workers[worker.Uuid] = worker

	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) UpdateDataNodes(ctx context.Context, nodesInfo *pbm.DataNodesInfo) (*pbm.Ack, error) {
	nodes := nodesInfo.GetNodes()
	datanodes := make(map[string]*DataNode)
	for _, node := range nodes {
		dataNode := &DataNode{
			Uuid:      node.GetUuid(),
			Filenames: node.GetFiles(),
		}
		datanodes[node.GetUuid()] = dataNode
	}
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.DFS.Nodes = datanodes
	return &pbm.Ack{
		Success: true,
	}, nil
}

func (s *Service) Trigger(ctx context.Context, taskRequest *pbm.TaskRequest) (*pbm.Ack, error) {

	return nil, nil
}
