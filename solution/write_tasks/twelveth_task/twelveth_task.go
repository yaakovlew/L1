package twelveth_task

import "fmt"

// Создание пустой струкутры, так как у мн-ва нет ключ - значение
type void struct{}

// инициализация пустой структуры
var empty void

func Twelveth(s []string) {
	//создание мн-ва
	set := make(map[string]void)
	//добавление элементов в мн-во
	for _, v := range s {
		set[v] = empty
	}
	//Результат
	fmt.Println(set)
}
