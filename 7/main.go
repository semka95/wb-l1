package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap[V any] struct {
	sync.RWMutex
	data map[string]V
}

func NewMap[V any]() ConcurrentMap[V] {
	return ConcurrentMap[V]{
		data: make(map[string]V),
	}
}

func (m *ConcurrentMap[V]) Get(key string) (V, bool) {
	// при rlock несколько горутин могут одновременно читать значение из мапы
	// но не могут изменять ее
	m.RLock()
	defer m.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *ConcurrentMap[V]) Set(key string, val V) {
	// при lock одна горутина может ее читать/изменять
	m.Lock()
	defer m.Unlock()
	m.data[key] = val
}

func (m *ConcurrentMap[V]) Has(key string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.data[key]
	return ok
}

func (m *ConcurrentMap[V]) Remove(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}

func main() {
	m := NewMap[int]()
	m.Set("test", 123)

	res, ok := m.Get("test")
	if ok {
		fmt.Printf("map[test]=%d\n", res)
	}

	ok = m.Has("foo")
	if !ok {
		fmt.Println("map does not have 'foo' key")
	}

	m.Remove("test")
	ok = m.Has("test")
	if !ok {
		fmt.Println("value deleted")
	}
}
