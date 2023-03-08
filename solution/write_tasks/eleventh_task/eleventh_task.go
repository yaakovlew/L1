package eleventh_task

import (
	"fmt"
)

type void struct{}

// инициализация пустой структуры
var empty void

func Intersection(s1, s2 map[string]void) map[string]void {
	set := make(map[string]void)
	for v, _ := range s1 {
		if _, exist := s2[v]; exist == true {
			set[v] = empty
		}
	}
	return set
}

func Eleventh(s1, s2 []string) {
	//создание первого мн-ва
	set1 := make(map[string]void)
	//добавление элементов в мн-во
	for _, v := range s1 {
		set1[v] = empty
	}
	//создание второго мн-ва
	set2 := make(map[string]void)
	//добавление элементов в мн-во
	for _, v := range s2 {
		set2[v] = empty
	}
	//Результат
	fmt.Println(Intersection(set1, set2))
}
