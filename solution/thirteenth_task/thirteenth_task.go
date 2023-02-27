package thirteenth_task

import "fmt"

func Thirteenth(a, b int) {
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
