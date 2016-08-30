package fastmap

import "sync"

type lockedVal struct {
	val interface{}
	l   *sync.RWMutex
}

type fastMap map[string]*lockedVal

func (f fastMap) Set(k string, v interface{}) {
	if v, ok := f[k]; !ok {
		f[k] = &lockedVal{val: v, l: &sync.RWMutex{}}
		return
	} else {
		f[k].l.Lock()
		defer f[k].l.Unlock()

		v.val = v
	}
}

func (f fastMap) Get(k string) (interface{}, bool) {
	f[k].l.RLock()
	defer f[k].l.RUnlock()

	if v, ok := f[k]; !ok {
		return nil, false
	} else {
		return v.val, true
	}
}

func (f fastMap) Del(k string) {
	if _, ok := f[k]; !ok {
		return
	}
	f[k].l.Lock()
	defer f[k].l.Unlock()
	delete(f, k)
}

func New() fastMap {
	return make(fastMap)
}
