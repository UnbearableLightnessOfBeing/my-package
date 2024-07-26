package myutils

import "sync"

type Set[T comparable] struct {
	data map[T]struct{}
	*sync.RWMutex
}

// NewSet creates a set data structure. It stores non-repeatable values in a thread safe map.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data:    map[T]struct{}{},
		RWMutex: &sync.RWMutex{},
	}
}

// Add stores values passed as parameters in a set.
func (s *Set[T]) Add(values... T) {
  s.Lock()
  defer s.Unlock()

  for _, v := range values {
    s.data[v] = struct{}{}
  }
} 

// Has indicates if a set containes a value passed as a parameter. 
//If a set containes a value passed as a parameter 
//it returns a truthy boolean value (true), (false) otherwise.
func (s *Set[T]) Has(v T) bool {
  s.RLock()
  defer s.RUnlock()

  _, ok := s.data[v]
  return ok
}

// GetElements returns a slice of values in a set.
func (s *Set[T]) GetElements() []T {
  s.RLock()
  defer s.RUnlock()

  var result []T
  for v := range s.data {
    result = append(result, v)
  }
  return result
}
