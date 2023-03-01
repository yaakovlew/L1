package third_task

import (
	"fmt"
	"sync"
)

func Third() {
	//создаение канала с 2-мя буфферами
	sum := make(chan int, 2)
	//запись в один из каналов 0
	sum <- 0
	w := sync.WaitGroup{}
	slice := []int{2, 4, 6, 8, 10}
	w.Add(len(slice))
	//Ф-ия возведения в квадрат
	squre := func(a int) {
		defer w.Done()
		sum <- (<-sum + a*a)
	}
	//Вызов ф-ий
	for _, v := range slice {
		go squre(v)
	}
	//ожидание завершения горутин
	w.Wait()
	//Результат
	fmt.Println(<-sum)
}
