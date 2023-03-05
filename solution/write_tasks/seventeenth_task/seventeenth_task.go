package seventeenth_task

import (
	"errors"
)

func Seventeen(mas []int, element int) (int, error) {
	l := 0
	r := len(mas) - 1
	for l <= r {
		var c int = (l + r) / 2
		if mas[c] == element {
			return c, nil
		} else if element > mas[c] {
			l = c + 1
		} else if element < mas[c] {
			r = c - 1
		} else {
			err := errors.New("Not Found")
			return 0, err
		}
	}
	err := errors.New("Not Found")
	return 0, err
}
