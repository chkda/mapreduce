package master

import (
	"context"
	"sync"

	pbm "github.com/chkda/mapreduce/rpc/master"
)

type Service struct {
	Workers map[string]*Worker
	Mu      sync.Mutex
}

type Worker struct {
	Uuid string
	IP   string
}

func New() *Service {
	return &Service{
		Workers: make(map[string]*Worker),
	}
}

func (s *Service) RegisterWorker(ctx context.Context, workerInfo *pbm.WorkerInfo) (*pbm.Ack, error) {
	worker := &Worker{
		Uuid: workerInfo.GetUuid(),
		IP:   workerInfo.GetIp(),
	}
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Workers[worker.Uuid] = worker

	return &pbm.Ack{
		Success: true,
	}, nil
}
