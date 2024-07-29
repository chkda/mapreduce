package worker

import (
	"context"
	"errors"
	"sync"

	pbw "github.com/chkda/mapreduce/rpc/worker"
)

var (
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

func (s *Service) GetReduceTask(taskId string) (*ReduceTask, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	task, ok := s.ReduceTasks[taskId]
	if !ok {
		return nil, ErrTaskIdDoesnotExist
	}
	return task, nil
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
