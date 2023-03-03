package tenth_task

import "math"

func Ten(t []float64) map[int][]float64 {
	group := make(map[int][]float64)
	for _, temp := range t {
		index := int(math.Floor(temp/10) * 10)
		group[index] = append(group[index], temp)
	}
	return group
}
