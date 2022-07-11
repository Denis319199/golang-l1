package task7

import "sync"

type AsyncMap[Key comparable, T any] interface {
}

type AsyncMapImpl[Key comparable, T any] struct {
	rwm sync.RWMutex
	m   map[Key]T
}

func (mp *AsyncMapImpl[Key, T]) Delete(key Key) {
	mp.rwm.Lock()
	delete(mp.m, key)
	mp.rwm.Unlock()
}

func (mp *AsyncMapImpl[Key, T]) Put(key Key, val *T) {
	mp.rwm.Lock()
	mp.m[key] = *val
	mp.rwm.Unlock()
}

func (mp *AsyncMapImpl[Key, T]) Get(key Key) (T, bool) {
	mp.rwm.RLock()
	val, ok := mp.m[key]
	mp.rwm.RUnlock()

	return val, ok
}
