package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (circle Circle) CalcPerimeter() float64 {
	lenCircle := 2 * math.Pi * circle.Radius
	return lenCircle
}

func (circle Circle) CalcArea() float64 {
	areaCircle := math.Pi * math.Pow(circle.Radius, 2)
	return areaCircle
}
