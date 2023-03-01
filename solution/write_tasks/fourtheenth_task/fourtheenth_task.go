package fourtheenth_task

import "fmt"

func Fourtheenth(def interface{}) {
	switch def.(type) {
	//если можно преобразовать в string, то ...
	case string:
		fmt.Println("It's string")
		//если можно преобразовать в int, то ...
	case int:
		fmt.Println("It's int")
		//если можно преобразовать в bool, то ...
	case bool:
		fmt.Println("It's bool")
		//иначе ...
	default:
		fmt.Println("It's chan")
	}
}
