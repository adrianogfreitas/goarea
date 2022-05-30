package area

import "math"

const Pi = 3.1416

func Circ(radius float64) float64 {
	return math.Pow(radius, 2) * Pi
}

func Rect(lenght, width float64) float64 {
	return lenght * width
}

func _RectTriangle(leg1, leg2 float64) float64 {
	return (leg1 * leg2) / 2
}
