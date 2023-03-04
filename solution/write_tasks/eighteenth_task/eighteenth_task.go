package eighteenth_task

import (
	"sync"
)

type Counter struct {
	m     sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.count++
}

func Eighteen() int {
	w := sync.WaitGroup{}
	c := new(Counter)
	c.count = 0
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func() {
			c.Increment()
			w.Done()
		}()
	}
	w.Wait()
	return c.count
}
