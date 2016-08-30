package lockedval

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("vim-go")
}

type lockedVal struct {
	val interface{}
	l   *sync.RWMutex
}

type lockedValMap map[string]*lockedVal

func (m lockedValMap) Set(k string, v interface{}) {
	if lv, ok := m[k]; !ok {
		m[k] = &lockedVal{val: v, l: &sync.RWMutex{}}
		return
	} else {
		m[k].l.Lock()
		defer m[k].l.Unlock()

		lv.val = v
	}

}

func (m lockedValMap) Get(k string) interface{} {
	m[k].l.RLock()
	defer m[k].l.RUnlock()

	return m[k]
}

func New() lockedValMap {
	return make(lockedValMap)
}
