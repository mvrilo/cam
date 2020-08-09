package cam

import (
	"sync"

	"gocv.io/x/gocv"
)

// Frame wraps the gocv.Mat and its index
type Frame struct {
	*Store
	Data  gocv.Mat
	Index int
}

type Store struct {
	data map[string]interface{}
	mu   *sync.RWMutex
}

func newStore() *Store {
	return &Store{
		data: make(map[string]interface{}),
		mu:   new(sync.RWMutex),
	}
}

func (s *Store) Get(name string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if val, found := s.data[name]; found {
		return val
	}
	return nil
}

func (s *Store) Set(name string, val interface{}) {
	s.mu.Lock()
	s.data[name] = val
	s.mu.Unlock()
}

func newFrame(index int) *Frame {
	return &Frame{
		Index: index,
		Data:  gocv.NewMat(),
		Store: newStore(),
	}
}
