package seventh_task

import (
	"fmt"
	"sync"
	"time"
)

// Структура map
type NewMap struct {
	mut sync.RWMutex
	m   map[string]int
}

// конструктор
func NewNewMap() *NewMap {
	return &NewMap{
		m: make(map[string]int),
	}
}

// Метод получения значения по ключу
func (c *NewMap) Load(key string) (int, bool) {
	val, ok := c.m[key]
	return val, ok
}

// Метод добавления значения по ключу
func (c *NewMap) Store(key string, value int, w *sync.WaitGroup) {
	defer w.Done()
	c.mut.RLock()
	defer c.mut.RUnlock()
	c.m[key] = value
	time.Sleep(3 * time.Second)
}

// Пример использования
func Seventh() {
	mapka := NewNewMap()
	w := &sync.WaitGroup{}
	w.Add(3)
	go mapka.Store("hello", 2, w)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(mapka.Load("hello"))
	go mapka.Store("world", 3, w)
	go mapka.Store("and you", 1, w)
	w.Wait()
	fmt.Println(mapka)
}
