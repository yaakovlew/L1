package third_task

import (
	"fmt"
	"sync"
)

func Third() {
	sum := make(chan int, 2)
	sum <- 0
	w := sync.WaitGroup{}
	slice := []int{2, 4, 6, 8, 10}
	w.Add(len(slice))
	squre := func(a int) {
		defer w.Done()
		sum <- (<-sum + a*a)
	}
	for i := range slice {
		go squre(i)
	}
	w.Wait()
	fmt.Println(<-sum)
}
