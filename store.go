package cam

import "sync"

// Storage interface
type Storage interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

// Store is a concurrency-safe map[string]interface{} implemntation
type Store struct {
	data map[string]interface{}
	mu   *sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]interface{}),
		mu:   new(sync.RWMutex),
	}
}

// Get returns a value as interface{} from its store
func (s *Store) Get(name string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if val, found := s.data[name]; found {
		return val
	}
	return nil
}

// Set writes a value on its store for a given name
func (s *Store) Set(name string, val interface{}) {
	s.mu.Lock()
	s.data[name] = val
	s.mu.Unlock()
}
