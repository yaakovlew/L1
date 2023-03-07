package seventh_task

import (
	"fmt"
	"sync"
)

// Структура map
type NewMap struct {
	mx sync.Mutex
	m  map[string]int
}

// конструктор
func NewNewMap() *NewMap {
	return &NewMap{
		m: make(map[string]int),
	}
}

// Метод получения значения по ключу
func (c *NewMap) Load(key string) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

// Метод добавления значения по ключу
func (c *NewMap) Store(key string, value int, w *sync.WaitGroup) {
	defer w.Done()
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

// Пример использования
func Seventh() {
	mapka := NewNewMap()
	w := &sync.WaitGroup{}
	w.Add(3)
	go mapka.Store("hello", 2, w)
	go mapka.Store("world", 3, w)
	go mapka.Store("and you", 1, w)
	w.Wait()
	fmt.Println(mapka)
}
