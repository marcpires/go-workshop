package inmem

import "sync"

type MemData struct {
	sync.Mutex // Solves type embed issue
	memo       map[string]int
}

func New() *MemData {
	return &MemData{memo: make(map[string]int)}
}

func (i *MemData) Get(k string) (int, bool) {
	i.Lock() // Acessing the Lock directly
	v, contains := i.memo[k]
	i.Unlock()
	return v, contains
}
