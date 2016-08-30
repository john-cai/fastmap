package maplocked

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("vim-go")
}

type lockedMap struct {
	m    map[string]interface{}
	lock *sync.RWMutex
}

func (m *lockedMap) Set(k string, v interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.m[k] = v
}

func (m *lockedMap) Get(k string) (interface{}, bool) {
	m.lock.RLock()
	m.lock.RUnlock()
	v, ok := m.m[k]
	return v, ok
}

func New() *lockedMap {
	return &lockedMap{
		m:    make(map[string]interface{}),
		lock: &sync.RWMutex{},
	}
}
