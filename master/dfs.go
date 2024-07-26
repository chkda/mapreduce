package master

import "errors"

var (
	ErrFileNotFound = errors.New("file not found")
)

type DataNode struct {
	Filenames []string
	Uuid      string
}

type HDFSReader interface {
	GetDataNodes(string) ([]*DataNode, error)
}

// Distributed File System
type DFS struct {
	FileChunks map[string][]*DataNode
}

func NewDFS() *DFS {
	return &DFS{
		FileChunks: make(map[string][]*DataNode),
	}
}

func (d *DFS) GetDataNodes(filename string) ([]*DataNode, error) {
	chunks, ok := d.FileChunks[filename]
	if !ok {
		return nil, ErrFileNotFound
	}
	return chunks, nil
}
