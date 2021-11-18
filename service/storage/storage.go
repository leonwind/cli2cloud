package storage

import (
	"sync"
)

type Storage struct {
	directoryPath string
	segmentSize   int
}

type segment struct {
	mu   sync.RWMutex
	size int
	ctr  int
}

func newSegment(filePath string, startCtr, size int) (*segment, error) {
	seg := new(segment)

	return seg, nil
}
