package twenty_fourth_task

import "math"

type Point struct {
	X float64
	Y float64
}

func TwentyFourth(point1 Point, point2 Point) (d float64) {
	d = math.Sqrt(math.Pow(point1.X-point2.X, 2) + math.Pow(point1.Y-point2.Y, 2))
	return d
}
