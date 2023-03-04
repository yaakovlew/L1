package twenty_first_task

import "fmt"

// структура 1
type FirstPackage struct{}

// структура 2
type SecondPackage struct{}

// функция структуры 2
func (s SecondPackage) SayHello2() string {
	return "Hello from second package"
}

// функция структуры 1
func (s FirstPackage) SayHello1() string {
	return "Hello from first package"
}

// 1 интерфейс
type First interface {
	SayHello1()
}

// 2 интерфейс
type Second interface {
	SayHello2()
}

// структура адаптер преобразующий интерфейс 1 в 2
type Adapter struct {
	First_pack FirstPackage
}

// функция соотвествующая 2 интерфейсу для 1 структуры
func (a Adapter) SayHello2() string {
	return a.First_pack.SayHello1()
}

// Проверка
func TwentyOne() {
	first := &FirstPackage{}
	second := &SecondPackage{}
	adap := &Adapter{First_pack: *first}
	fmt.Println(second.SayHello2())
	fmt.Println(adap.SayHello2())
}
