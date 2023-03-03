package nineth_task

import (
	"fmt"
	"sync"
)

func Nine(mas []int) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, v := range mas {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for _, v := range mas {
			ch2 <- v * 2
		}
		close(ch2)
	}()
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		for num := range ch2 {
			fmt.Println(num)
		}
		w.Done()
	}()
	w.Wait()
}
