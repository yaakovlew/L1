package sixteenth_task

import (
	"fmt"
	"sort"
)

func Sixteenth(mas []int) {
	sort.Slice(mas, func(i, j int) bool {
		return mas[i] < mas[j]
	})
	fmt.Println(mas)
}
