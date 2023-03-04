package fifth_task

import (
	"fmt"
	"time"
)

func writeInChanel(ch chan int) {
	i := 0
	for {
		i++
		time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
}

func readFromChanel(ch chan int) {
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Chanel close")
}

func Five(sec int) {
	ch := make(chan int)
	go writeInChanel(ch)
	go readFromChanel(ch)
	time.Sleep(time.Duration(sec) * time.Second)
}
