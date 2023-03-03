package sixth_task

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func Six() {
	Six4()
}

// Первый способ
func Six1() {
	ctx, cancel := context.WithCancel(context.Background())
	someFunc1 := func(ctx context.Context) {
		time.Sleep(10 * time.Second)
		fmt.Println("Done")
	}
	go someFunc1(ctx)
	time.Sleep(3 * time.Second)
	cancel()
}

// Второй способ
func Six2() {
	someFunc2 := func(stop chan bool) {
		for {
			select {
			case <-stop:
				fmt.Println("someFunc2 stopped")
				return
			default:
				fmt.Println("someFunc2 running")
				time.Sleep(1 * time.Second)
			}
		}
	}
	stop := make(chan bool)
	go someFunc2(stop)

	time.Sleep(5 * time.Second)
	stop <- true
}

// 3 способ
func Six3() {
	someFunc3 := func(ch chan string) {
		for {
			select {
			case <-ch:
				fmt.Println("someFunc3 stopped")
				return
			default:
				fmt.Println("someFunc3 running")
				time.Sleep(1 * time.Second)
			}
		}
	}
	ch := make(chan string)
	go someFunc3(ch)
	time.Sleep(5 * time.Second)
	close(ch)
	time.Sleep(2 * time.Second)
}

// 4 способ
func Six4() {
	someFunc4 := func() {
		//принудительное завершение
		runtime.Goexit()
		time.Sleep(10 * time.Second)
		fmt.Println("Done")
	}
	go someFunc4()
}
