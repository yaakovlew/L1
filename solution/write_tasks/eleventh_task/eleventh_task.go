package eleventh_task

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func Eleven() {
	//Создание и инициализация первого мн-ва
	required := mapset.NewSet()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")
	//Создание и инициализация второго мн-ва
	sciences := mapset.NewSet()
	sciences.Add("biology")
	sciences.Add("chemistry")
	//Пересечение этих мн-в
	count := required.Intersect(sciences)
	fmt.Println(count)
}
