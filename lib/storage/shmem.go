package storage

import (
	"log"
	"os"
	"path"
)

const ARRAY_SIZE = 4

type NodeIndex uint32
type Epoch uint32

type Storage struct {
	// The directory where the files are stored
	root string

	epoch Epoch

	nodes typedData[Node, NodeIndex]
}

// TODO(eric): Make these a linked list for concurrent executions to work
// Stored in mmap'd file
type Node struct {
	validated Epoch
	numDeps   uint8
	deps      [4]NodeIndex
}

func Stuff(path string) {
	storage := Open(path)
	defer storage.Close()

	MakeChain(storage)

	log.Printf("Have %d nodes", len(storage.nodes.d))

	log.Printf("Node 1: %+v", storage.nodes.Get(1))

}

func MakeChain(s *Storage) {
	for i := NodeIndex(0); i < ARRAY_SIZE-1; i++ {
		n := s.nodes.Get(i)
		n.numDeps = 1
		n.deps[0] = i + 1
	}
}

func Open(root string) *Storage {
	// Create the directory if needed
	if err := os.MkdirAll(root, 0777); err != nil {
		log.Fatal(err)
	}

	return &Storage{
		root:  root,
		epoch: 1,
		nodes: MmapTyped[Node, NodeIndex](path.Join(root, "file1.bin")),
	}

}

func (s *Storage) Close() {
	s.nodes.file.Close()
}
