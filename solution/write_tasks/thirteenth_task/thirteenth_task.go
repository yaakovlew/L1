package thirteenth_task

import "fmt"

func Thirteenth(a, b int) {
	fmt.Println(a, b)
	//замена а на b, a b на а
	a, b = b, a
	fmt.Println(a, b)
}
