package twenty_fourth_task

import "math"

// Структура Point
type Point struct {
	X float64
	Y float64
}

func TwentyFourth(point1 Point, point2 Point) (d float64) {
	//формула вычисления растояние между 2-мя точками в 2-ухмерном пространстве
	d = math.Sqrt(math.Pow(point1.X-point2.X, 2) + math.Pow(point1.Y-point2.Y, 2))
	return d
}
