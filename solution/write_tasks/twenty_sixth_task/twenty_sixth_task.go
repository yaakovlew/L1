package twenty_sixth_task

import (
	"strings"
)

type void struct{}

var empty void

func TwentySixth(str string) bool {
	str = strings.ToLower(str)
	check := make(map[int32]void)
	for _, v := range str {
		check[v] = empty
	}
	if len(str) == len(check) {
		return true
	} else {
		return false
	}
}
