package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rec Rectangle) CalcPerimeter() float64 {
	perim := 2 * (rec.Height + rec.Weight)
	return perim
}

func (rec Rectangle) CalcArea() float64 {
	area := rec.Height * rec.Weight
	return area
}
