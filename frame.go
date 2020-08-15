package cam

import (
	"sync"

	"gocv.io/x/gocv"
)

// Frame represents a camera frame wrapped up as gocv.Mat along with its index and a key-value store
type Frame struct {
	*Store
	Data  gocv.Mat
	Index int
}

// Store is a concurrency-safe map[string]interface{} implemntation
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

// Get gets a value as interface{} from its store
func (s *Store) Get(name string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if val, found := s.data[name]; found {
		return val
	}
	return nil
}

// Set sets a value on its store for a given name
func (s *Store) Set(name string, val interface{}) {
	s.mu.Lock()
	s.data[name] = val
	s.mu.Unlock()
}

func newFrame() *Frame {
	return &Frame{
		Index: 0,
		Data:  gocv.NewMat(),
		Store: newStore(),
	}
}
