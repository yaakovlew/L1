package twelveth_task

import "fmt"

type void struct{}

var empty void

func Twelveth(s []string) {
	set := make(map[string]void)
	for _, v := range s {
		set[v] = empty
	}
	fmt.Println(set)
}
