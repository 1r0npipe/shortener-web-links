package model

import (
	"sync"
)

type Item struct {
	Ref   uint
	Short string
	Long  string
	Ttl   int64
}
type Info struct {
	mu   sync.RWMutex
	link map[string]*Item
}

func (i *Info) Get(key string) *Item {
	i.mu.RLock()
	if val, ok := i.link[key]; ok {
		return val
	}
	i.mu.RUnlock()
	return nil
}

func (i *Info) Put(key string, item *Item) {
	i.mu.Lock()
	i.link[key] = item
	i.mu.Unlock()
}

func (i *Info) Delete(key string) bool {
	if _, ok := i.link[key]; ok {
		i.mu.Lock()
		delete(i.link, key)
		i.mu.Unlock()
		return true
	}
	return false
}
