package seventh_task

import (
	"fmt"
	"sync"
)

type NewMap struct {
	mx sync.Mutex
	m  map[string]int
}

func NewNewMap() *NewMap {
	return &NewMap{
		m: make(map[string]int),
	}
}

func (c *NewMap) Load(key string) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *NewMap) Store(key string, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func Seventh() {
	mapka := NewNewMap()
	go mapka.Store("hello", 2)
	go mapka.Store("world", 3)
	go mapka.Store("and you", 1)
	fmt.Println(mapka)
}
