package master

import (
	"testing"
)

func TestNewDFS(t *testing.T) {
	dfs := NewDFS()
	if dfs == nil {
		t.Fatal("NewDFS() returned nil")
	}
	if dfs.FileChunks == nil {
		t.Error("FileChunks map not initialized")
	}
	if len(dfs.FileChunks) != 0 {
		t.Errorf("Expected empty FileChunks, got %d entries", len(dfs.FileChunks))
	}
}

func TestDFS_GetDataNodes_Success(t *testing.T) {
	dfs := NewDFS()
	
	// Setup test data
	expectedNodes := []*DataNode{
		{
			Uuid:      "worker-1",
			Filenames: []string{"chunk-1.txt", "chunk-2.txt"},
		},
		{
			Uuid:      "worker-2",
			Filenames: []string{"chunk-3.txt"},
		},
	}
	
	dfs.FileChunks["test-file.txt"] = expectedNodes
	
	// Test retrieval
	nodes, err := dfs.GetDataNodes("test-file.txt")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	
	if len(nodes) != len(expectedNodes) {
		t.Errorf("Expected %d nodes, got %d", len(expectedNodes), len(nodes))
	}
	
	// Verify first node
	if nodes[0].Uuid != "worker-1" {
		t.Errorf("Expected first node UUID 'worker-1', got '%s'", nodes[0].Uuid)
	}
	if len(nodes[0].Filenames) != 2 {
		t.Errorf("Expected first node to have 2 filenames, got %d", len(nodes[0].Filenames))
	}
}

func TestDFS_GetDataNodes_FileNotFound(t *testing.T) {
	dfs := NewDFS()
	
	// Try to get non-existent file
	nodes, err := dfs.GetDataNodes("non-existent.txt")
	
	if err != ErrFileNotFound {
		t.Errorf("Expected ErrFileNotFound, got: %v", err)
	}
	
	if nodes != nil {
		t.Errorf("Expected nil nodes, got: %v", nodes)
	}
}

func TestDFS_MultipleFiles(t *testing.T) {
	dfs := NewDFS()
	
	// Add multiple files
	dfs.FileChunks["file1.txt"] = []*DataNode{
		{Uuid: "worker-1", Filenames: []string{"f1-chunk1.txt"}},
	}
	dfs.FileChunks["file2.txt"] = []*DataNode{
		{Uuid: "worker-2", Filenames: []string{"f2-chunk1.txt"}},
	}
	
	// Retrieve file1
	nodes1, err := dfs.GetDataNodes("file1.txt")
	if err != nil {
		t.Fatalf("Error getting file1: %v", err)
	}
	if nodes1[0].Uuid != "worker-1" {
		t.Errorf("Expected worker-1 for file1, got %s", nodes1[0].Uuid)
	}
	
	// Retrieve file2
	nodes2, err := dfs.GetDataNodes("file2.txt")
	if err != nil {
		t.Fatalf("Error getting file2: %v", err)
	}
	if nodes2[0].Uuid != "worker-2" {
		t.Errorf("Expected worker-2 for file2, got %s", nodes2[0].Uuid)
	}
}

func TestDFS_EmptyFilenames(t *testing.T) {
	dfs := NewDFS()
	
	// Node with no filenames
	dfs.FileChunks["empty.txt"] = []*DataNode{
		{Uuid: "worker-1", Filenames: []string{}},
	}
	
	nodes, err := dfs.GetDataNodes("empty.txt")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	
	if len(nodes[0].Filenames) != 0 {
		t.Errorf("Expected 0 filenames, got %d", len(nodes[0].Filenames))
	}
}
