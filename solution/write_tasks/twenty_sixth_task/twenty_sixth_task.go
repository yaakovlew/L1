package twenty_sixth_task

import (
	"strings"
)

type void struct{}

var empty void

func TwentySixth(str string) bool {
	//изменение всех элементов в строке в нижний регистр
	str = strings.ToLower(str)
	//создание мн-ва
	check := make(map[int32]void)
	//добавление символов в мн-во
	for _, v := range str {
		check[v] = empty
	}
	//если кол-во элементов в мн-ве равно кол-ву элементов в строке, то все элементы уникальны
	if len(str) == len(check) {
		return true
	} else {
		return false
	}
}
