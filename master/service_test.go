package master

import (
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"single", "a"},
		{"word", "hello"},
		{"sentence", "hello world"},
		{"numbers", "12345"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash1 := hash(tt.input)
			hash2 := hash(tt.input)
			
			// Same input should produce same hash
			if hash1 != hash2 {
				t.Errorf("Hash not consistent: %d != %d", hash1, hash2)
			}
		})
	}
}

func TestHash_Distribution(t *testing.T) {
	// Test that different inputs produce different hashes
	inputs := []string{"key1", "key2", "key3", "key4", "key5"}
	hashes := make(map[uint32]bool)
	
	for _, input := range inputs {
		h := hash(input)
		hashes[h] = true
	}
	
	// We expect most inputs to have unique hashes
	if len(hashes) < 4 {
		t.Errorf("Expected at least 4 unique hashes, got %d", len(hashes))
	}
}

func TestService_CalculateReducerIdFromFilePartition(t *testing.T) {
	s := &Service{NumReduce: 5}
	
	tests := []struct {
		name     string
		filename string
		want     uint32
		wantErr  bool
	}{
		{
			name:     "valid partition file",
			filename: "mr-abc123-3.txt",
			want:     3,
			wantErr:  false,
		},
		{
			name:     "partition 0",
			filename: "mr-xyz789-0.txt",
			want:     0,
			wantErr:  false,
		},
		{
			name:     "wrong format - missing mr prefix",
			filename: "task-abc-1.txt",
			want:     0,
			wantErr:  true,
		},
		{
			name:     "wrong format - too few parts",
			filename: "mr-abc.txt",
			want:     0,
			wantErr:  true,
		},
		{
			name:     "invalid reducer id - not a number",
			filename: "mr-abc-xyz.txt",
			want:     0,
			wantErr:  true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.calculateReducerIdFromFilePartition(tt.filename)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}
			
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			
			if got != tt.want {
				t.Errorf("Expected reducer ID %d, got %d", tt.want, got)
			}
		})
	}
}

func TestService_GetReducerIdToWorker(t *testing.T) {
	s := &Service{
		NumReduce: 3,
		Workers: map[string]*Worker{
			"worker-1": {Uuid: "worker-1"},
			"worker-2": {Uuid: "worker-2"},
			"worker-3": {Uuid: "worker-3"},
		},
	}
	
	reducerMap := s.getReducerIdToWorker()
	
	// Should have entries (might be less than NumReduce due to collisions)
	if len(reducerMap) == 0 {
		t.Error("Reducer map is empty")
	}
	
	// All reducer IDs should be < NumReduce
	for reducerId := range reducerMap {
		if reducerId >= uint32(s.NumReduce) {
			t.Errorf("Reducer ID %d is >= NumReduce %d", reducerId, s.NumReduce)
		}
	}
	
	// All workers should be valid
	for _, workerId := range reducerMap {
		if _, exists := s.Workers[workerId]; !exists {
			t.Errorf("Worker ID %s not found in worker map", workerId)
		}
	}
}

func TestService_AddWorker(t *testing.T) {
	s := &Service{
		Workers: make(map[string]*Worker),
	}
	
	worker := &Worker{
		Uuid: "worker-123",
		IP:   "localhost:5001",
	}
	
	s.AddWorker("worker-123", worker)
	
	if len(s.Workers) != 1 {
		t.Errorf("Expected 1 worker, got %d", len(s.Workers))
	}
	
	gotWorker, err := s.GetWorker("worker-123")
	if err != nil {
		t.Fatalf("Error getting worker: %v", err)
	}
	
	if gotWorker.Uuid != "worker-123" {
		t.Errorf("Expected worker UUID 'worker-123', got '%s'", gotWorker.Uuid)
	}
}

func TestService_GetWorker_NotFound(t *testing.T) {
	s := &Service{
		Workers: make(map[string]*Worker),
	}
	
	_, err := s.GetWorker("non-existent")
	if err != ErrWorkerIdDoesnotExist {
		t.Errorf("Expected ErrWorkerIdDoesnotExist, got: %v", err)
	}
}

func TestService_TaskCounters(t *testing.T) {
	s := &Service{
		ActiveMapTasks:    0,
		ActiveReduceTasks: 0,
	}
	
	// Test map task counters
	s.IncActiveMapTasks()
	s.IncActiveMapTasks()
	if s.ActiveMapTasks != 2 {
		t.Errorf("Expected 2 active map tasks, got %d", s.ActiveMapTasks)
	}
	
	s.DecActiveMapTasks()
	if s.ActiveMapTasks != 1 {
		t.Errorf("Expected 1 active map task, got %d", s.ActiveMapTasks)
	}
	
	// Test reduce task counters
	s.IncActiveReduceTasks()
	s.IncActiveReduceTasks()
	s.IncActiveReduceTasks()
	if s.ActiveReduceTasks != 3 {
		t.Errorf("Expected 3 active reduce tasks, got %d", s.ActiveReduceTasks)
	}
	
	s.DecActiveReduceTasks()
	if s.ActiveReduceTasks != 2 {
		t.Errorf("Expected 2 active reduce tasks, got %d", s.ActiveReduceTasks)
	}
}

func TestService_DeadWorkers(t *testing.T) {
	s := &Service{
		DeadWorkers:          make(map[string]bool),
		DeadWorkersThreshold: 2,
	}
	
	// Initially no dead workers
	if s.checkIfDeadWorkersExceedsThreshold() {
		t.Error("Should not exceed threshold with 0 dead workers")
	}
	
	// Add one dead worker
	s.AddDeadWorkers("worker-1")
	if s.checkIfDeadWorkersExceedsThreshold() {
		t.Error("Should not exceed threshold with 1 dead worker (threshold=2)")
	}
	
	// Add second dead worker - should exceed
	s.AddDeadWorkers("worker-2")
	if !s.checkIfDeadWorkersExceedsThreshold() {
		t.Error("Should exceed threshold with 2 dead workers (threshold=2)")
	}
	
	// Remove one dead worker
	s.RemoveDeadWorkers("worker-1")
	if s.checkIfDeadWorkersExceedsThreshold() {
		t.Error("Should not exceed threshold after removing one worker")
	}
}

func TestNewService(t *testing.T) {
	cfg := &Config{
		GRPCPort:             "4000",
		Filename:             "test-input.txt",
		NumReduce:            5,
		DeadWorkersThreshold: 3,
	}
	
	s := New(cfg)
	
	if s == nil {
		t.Fatal("New() returned nil")
	}
	
	if s.Filename != "test-input.txt" {
		t.Errorf("Expected filename 'test-input.txt', got '%s'", s.Filename)
	}
	
	if s.NumReduce != 5 {
		t.Errorf("Expected NumReduce 5, got %d", s.NumReduce)
	}
	
	if s.DeadWorkersThreshold != 3 {
		t.Errorf("Expected DeadWorkersThreshold 3, got %d", s.DeadWorkersThreshold)
	}
	
	if s.Workers == nil {
		t.Error("Workers map not initialized")
	}
	
	if s.DFS == nil {
		t.Error("DFS not initialized")
	}
	
	if s.DeadWorkers == nil {
		t.Error("DeadWorkers map not initialized")
	}
}
