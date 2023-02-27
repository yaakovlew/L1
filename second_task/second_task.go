package main

import (
	"fmt"
	"sync"
)

func main() {
	w := sync.WaitGroup{}
	slice := []int{2, 4, 6, 8, 10}
	w.Add(len(slice))
	squre := func(a int) {
		w.Done()
		fmt.Println(a*a, "is square of", a)
	}
	for i := range slice {
		go squre(i)
	}
	w.Wait()
}
