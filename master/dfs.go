package master

type DataNode struct {
	Filenames []string
	Uuid      string
}

type HDFSNodeReader interface {
	GetDataNodes(string) (*DataNode, error)
}

// Distributed File System
type DFS struct {
	Nodes map[string]*DataNode
}

func NewDFS() *DFS {
	return &DFS{
		Nodes: make(map[string]*DataNode),
	}
}
