package fourth_task

import (
	"fmt"
	"os"
	"os/signal"
)

// главный поток
func writeToChannel(ch chan<- string) {
	for {
		var input string
		fmt.Scanln(&input)
		ch <- input
	}
}

// воркеры
func worker(id int, ch <-chan string) {
	for data := range ch {
		fmt.Printf("Worker %d: %s\n", id, data)
	}
}

// запуск воркеров и главного потока
func Fourth(numWorkers int) {
	ch := make(chan string)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}
	go writeToChannel(ch)

	//сигнал прерывания
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	close(ch)
}
