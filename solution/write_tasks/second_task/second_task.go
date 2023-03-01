package second_task

import (
	"fmt"
	"sync"
)

func Second() {
	w := sync.WaitGroup{}
	slice := []int{2, 4, 6, 8, 10}
	// должны дождаться работу всех горутин, а их = кол-ву элементов
	w.Add(len(slice))
	//функция возведения в квадрат
	squre := func(a int) {
		defer w.Done()
		fmt.Println(a*a, "is square of", a)
	}
	//параллельный вызов функций
	for i := range slice {
		go squre(i)
	}
	//ожидание завершения горутин
	w.Wait()
}
