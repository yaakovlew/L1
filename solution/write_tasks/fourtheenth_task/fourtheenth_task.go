package fourtheenth_task

import "fmt"

func Fourtheenth(def interface{}) {
	switch def.(type) {
	case string:
		fmt.Println("It's string")
	case int:
		fmt.Println("It's int")
	case bool:
		fmt.Println("It's bool")
	default:
		fmt.Println("It's chan")
	}
}
