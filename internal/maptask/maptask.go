package maptask

import (
	"strings"
	"sync"

	"github.com/google/uuid"
)

type MapTaskStatus int

const (
	IDLE MapTaskStatus = iota
	INPROGRESS
	COMPLETE
	FAILED
)

type Option func(s *Task)

type Task struct {
	mu          sync.RWMutex
	taskId      string
	workerId    string
	taskFile    string
	numReduce   int
	taskStatus  MapTaskStatus
	outputFiles []string
}

func New(opts ...Option) *Task {
	task := &Task{
		taskId:      strings.ReplaceAll(uuid.NewString(), "-", ""),
		taskStatus:  IDLE,
		outputFiles: make([]string, 0),
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

func WithTaskFile(fileName string) Option {
	return func(s *Task) {
		s.taskFile = fileName
	}
}

func WithNumReduce(numReduce int) Option {
	return func(s *Task) {
		s.numReduce = numReduce
	}
}

func (s *Task) GetTaskId() string {
	return s.taskId
}

func (s *Task) GetTaskFile() string {
	return s.taskFile
}

func (s *Task) GetWorkerId() string {
	return s.workerId
}

func (s *Task) GetNumReduce() int {
	return s.numReduce
}

func (s *Task) GetTaskStatus() MapTaskStatus {
	s.mu.RLock()
	defer s.mu.Unlock()
	return s.taskStatus
}

func (s *Task) SetTaskStatus(status MapTaskStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.taskStatus = status
}

func (s *Task) GetOutputFiles() []string {
	s.mu.RLock()
	defer s.mu.Unlock()
	return s.outputFiles
}

func (s *Task) SetOutputFiles(files []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.outputFiles = files
}
