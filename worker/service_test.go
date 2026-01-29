package worker

import (
	"bytes"
	"strings"
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

func TestHash_Consistency(t *testing.T) {
	// Test that hash function is consistent with master
	// Both should use FNV-1a 32-bit hash
	testCases := []string{
		"hello",
		"world",
		"mapreduce",
		"test123",
	}
	
	for _, tc := range testCases {
		h := hash(tc)
		if h == 0 {
			t.Errorf("Hash of '%s' should not be 0", tc)
		}
	}
}

func TestService_Map(t *testing.T) {
	s := &Service{}
	
	input := "hello world hello test world"
	emittedPairs := make(map[string]int)
	
	emitFunc := func(key, value string) {
		emittedPairs[key]++
	}
	
	s.Map("test.txt", input, emitFunc)
	
	// Should emit 5 pairs (one for each word)
	totalEmissions := 0
	for _, count := range emittedPairs {
		totalEmissions += count
	}
	
	if totalEmissions != 5 {
		t.Errorf("Expected 5 emissions, got %d", totalEmissions)
	}
	
	// Check individual words
	expectedWords := map[string]int{
		"hello": 2,
		"world": 2,
		"test":  1,
	}
	
	for word, expectedCount := range expectedWords {
		if emittedPairs[word] != expectedCount {
			t.Errorf("Word '%s': expected %d emissions, got %d", 
				word, expectedCount, emittedPairs[word])
		}
	}
}

func TestService_Reduce(t *testing.T) {
	s := &Service{}
	
	tests := []struct {
		name   string
		key    string
		values []string
		want   string
	}{
		{
			name:   "single value",
			key:    "hello",
			values: []string{"1"},
			want:   "1",
		},
		{
			name:   "multiple values",
			key:    "world",
			values: []string{"1", "1", "1"},
			want:   "3",
		},
		{
			name:   "empty values",
			key:    "test",
			values: []string{},
			want:   "0",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Reduce(tt.key, tt.values)
			if got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ProcessIntermediateData(t *testing.T) {
	s := &Service{}
	
	// Simulate intermediate data in tab-separated format
	data := []byte("hello\t1\nworld\t1\nhello\t1\ntest\t1\nworld\t1\n")
	kvMap := make(map[string][]string)
	
	s.processIntermediateData(data, kvMap)
	
	// Check "hello" appears twice
	if len(kvMap["hello"]) != 2 {
		t.Errorf("Expected 'hello' to have 2 values, got %d", len(kvMap["hello"]))
	}
	
	// Check "world" appears twice
	if len(kvMap["world"]) != 2 {
		t.Errorf("Expected 'world' to have 2 values, got %d", len(kvMap["world"]))
	}
	
	// Check "test" appears once
	if len(kvMap["test"]) != 1 {
		t.Errorf("Expected 'test' to have 1 value, got %d", len(kvMap["test"]))
	}
}

func TestService_ProcessIntermediateData_InvalidFormat(t *testing.T) {
	s := &Service{}
	
	// Data with invalid format (no tabs)
	data := []byte("hello 1\nworld 1\n")
	kvMap := make(map[string][]string)
	
	s.processIntermediateData(data, kvMap)
	
	// Should not process invalid lines
	if len(kvMap) != 0 {
		t.Errorf("Expected empty kvMap for invalid data, got %d entries", len(kvMap))
	}
}

func TestService_ProcessIntermediateData_EmptyData(t *testing.T) {
	s := &Service{}
	
	data := []byte("")
	kvMap := make(map[string][]string)
	
	s.processIntermediateData(data, kvMap)
	
	if len(kvMap) != 0 {
		t.Errorf("Expected empty kvMap for empty data, got %d entries", len(kvMap))
	}
}

func TestService_AddMapTask(t *testing.T) {
	s := &Service{
		MapTasks: make(map[string]*MapTask),
	}
	
	task := NewMapTask("file.txt", "task-1", 3)
	s.AddMapTask("task-1", task)
	
	if len(s.MapTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(s.MapTasks))
	}
	
	gotTask, err := s.GetMapTask("task-1")
	if err != nil {
		t.Fatalf("Error getting task: %v", err)
	}
	
	if gotTask.GetId() != "task-1" {
		t.Errorf("Expected task ID 'task-1', got '%s'", gotTask.GetId())
	}
}

func TestService_GetMapTask_NotFound(t *testing.T) {
	s := &Service{
		MapTasks: make(map[string]*MapTask),
	}
	
	_, err := s.GetMapTask("non-existent")
	if err != ErrTaskIdDoesnotExist {
		t.Errorf("Expected ErrTaskIdDoesnotExist, got: %v", err)
	}
}

func TestService_AddReduceTask(t *testing.T) {
	s := &Service{
		ReduceTasks: make(map[string]*ReduceTask),
	}
	
	taskFiles := []*ReduceDataNodeInfo{
		{Filename: "test.txt", NodeIP: "localhost:5001"},
	}
	task := NewReduceTask(taskFiles, "task-1")
	s.AddReduceTask("task-1", task)
	
	if len(s.ReduceTasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(s.ReduceTasks))
	}
	
	gotTask, err := s.GetReduceTask("task-1")
	if err != nil {
		t.Fatalf("Error getting task: %v", err)
	}
	
	if gotTask.GetId() != "task-1" {
		t.Errorf("Expected task ID 'task-1', got '%s'", gotTask.GetId())
	}
}

func TestService_Map_EmptyInput(t *testing.T) {
	s := &Service{}
	
	input := ""
	emittedCount := 0
	
	emitFunc := func(key, value string) {
		emittedCount++
	}
	
	s.Map("test.txt", input, emitFunc)
	
	if emittedCount != 0 {
		t.Errorf("Expected 0 emissions for empty input, got %d", emittedCount)
	}
}

func TestService_Map_Whitespace(t *testing.T) {
	s := &Service{}
	
	input := "   hello    world   "
	words := make([]string, 0)
	
	emitFunc := func(key, value string) {
		words = append(words, key)
	}
	
	s.Map("test.txt", input, emitFunc)
	
	if len(words) != 2 {
		t.Errorf("Expected 2 words, got %d", len(words))
	}
	
	if words[0] != "hello" || words[1] != "world" {
		t.Errorf("Expected [hello, world], got %v", words)
	}
}

func TestPartitioning(t *testing.T) {
	// Test that the same key always goes to the same partition
	numReduce := 5
	key := "test-key"
	
	partition1 := hash(key) % uint32(numReduce)
	partition2 := hash(key) % uint32(numReduce)
	
	if partition1 != partition2 {
		t.Errorf("Partitioning not consistent: %d != %d", partition1, partition2)
	}
	
	// Partition should be within valid range
	if partition1 >= uint32(numReduce) {
		t.Errorf("Partition %d out of range [0, %d)", partition1, numReduce)
	}
}

func TestService_ProcessIntermediateData_MixedContent(t *testing.T) {
	s := &Service{}
	
	var buf bytes.Buffer
	buf.WriteString("key1\tvalue1\n")
	buf.WriteString("key2\tvalue2\n")
	buf.WriteString("key1\tvalue3\n")
	buf.WriteString("invalid line\n")  // No tab
	buf.WriteString("key3\tvalue4\n")
	
	kvMap := make(map[string][]string)
	s.processIntermediateData(buf.Bytes(), kvMap)
	
	// Should have 3 keys (invalid line skipped)
	if len(kvMap) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(kvMap))
	}
	
	// key1 should have 2 values
	if len(kvMap["key1"]) != 2 {
		t.Errorf("Expected key1 to have 2 values, got %d", len(kvMap["key1"]))
	}
}

func TestMapWordCount_Integration(t *testing.T) {
	// Integration test: Map -> Partition -> Reduce
	s := &Service{}
	
	// Simulate map phase
	input := "hello world hello"
	kvPairs := make(map[string][]string)
	
	emitFunc := func(key, value string) {
		kvPairs[key] = append(kvPairs[key], value)
	}
	
	s.Map("input.txt", input, emitFunc)
	
	// Simulate reduce phase
	for key, values := range kvPairs {
		result := s.Reduce(key, values)
		
		if key == "hello" && result != "2" {
			t.Errorf("Expected 'hello' count to be 2, got %s", result)
		}
		if key == "world" && result != "1" {
			t.Errorf("Expected 'world' count to be 1, got %s", result)
		}
	}
}

func BenchmarkHash(b *testing.B) {
	input := "benchmark-key-test-string"
	for i := 0; i < b.N; i++ {
		_ = hash(input)
	}
}

func BenchmarkMap(b *testing.B) {
	s := &Service{}
	input := strings.Repeat("hello world ", 100)
	emitFunc := func(key, value string) {}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Map("test.txt", input, emitFunc)
	}
}
