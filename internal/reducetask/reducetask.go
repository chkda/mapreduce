package reducetask

import (
	"strings"
	"sync"

	"github.com/chkda/mapreduce/internal/datanodes"
	"github.com/chkda/mapreduce/internal/status"
	"github.com/google/uuid"
)

type Task struct {
	taskId     string
	mu         sync.RWMutex
	taskFiles  []*datanodes.ReduceDataNodeInfo
	workerId   string
	taskStatus status.TaskStatus
	outputFile string
}

type Option func(s *Task)

func New(opts ...Option) *Task {
	task := &Task{
		taskId:     strings.ReplaceAll(uuid.NewString(), "-", ""),
		taskStatus: status.IDLE,
	}

	for _, opt := range opts {
		opt(task)
	}

	return task
}

func WithTaskId(id string) Option {
	return func(s *Task) {
		s.taskId = id
	}
}

func WithWorkerId(id string) Option {
	return func(s *Task) {
		s.workerId = id
	}
}

func WithTaskFiles(files []*datanodes.ReduceDataNodeInfo) Option {
	return func(s *Task) {
		s.taskFiles = files
	}
}

func (s *Task) GetTaskId() string {
	return s.taskId
}

func (s *Task) GetTaskFiles() []*datanodes.ReduceDataNodeInfo {
	return s.taskFiles
}

func (s *Task) GetWorkerId() string {
	return s.workerId
}

func (s *Task) GetTaskStatus() status.TaskStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.taskStatus
}

func (s *Task) SetTaskStatus(status status.TaskStatus) {
	s.mu.Lock()
	s.taskStatus = status
	s.mu.Unlock()
}

func (s *Task) GetOutputFile() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.outputFile
}

func (s *Task) SetOutputFile(files string) {
	s.mu.Lock()
	s.outputFile = files
	s.mu.Unlock()
}
