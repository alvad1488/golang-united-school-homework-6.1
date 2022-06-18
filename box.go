package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	len := len(b.shapes)
	cap := b.shapesCapacity

	if len < cap {
		b.shapes = append(b.shapes, shape)
		return nil
	}

	return fmt.Errorf("you cannot add more shapes than capacity allow: len= %d, cap= %d", len, cap)
}

//Remove by index function
func (b *box) remove(i int) error {
	len := len(b.shapes)
	if len == 0 {
		return fmt.Errorf("zero length slice")
	}

	b.shapes[i] = b.shapes[len-1]
	b.shapes[len-1] = nil
	b.shapes = b.shapes[:len-1]
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	len := len(b.shapes)
	if i > len-1 {
		return nil, fmt.Errorf("index out of range")
	}

	data := b.shapes[i]
	if data == nil {
		return data, fmt.Errorf("by index %d shape not found", i)
	}

	return data, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	len := len(b.shapes)
	if i > len-1 {
		return nil, fmt.Errorf("index out of range")
	}

	data := b.shapes[i]
	if data == nil {
		return data, fmt.Errorf("by index %d shape not found", i)
	}

	b.remove(i)
	return data, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	len := len(b.shapes)
	if i > len-1 {
		return nil, fmt.Errorf("index out of range")
	}

	data := b.shapes[i]
	if data == nil {
		return data, fmt.Errorf("by index %d shape not found", i)
	}

	b.shapes[i] = shape
	return data, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64 = 0
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64 = 0
	for _, shape := range b.shapes {
		sumArea += shape.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	cntCircle := 0
	for i := 0; i < len(b.shapes); {
		shape := b.shapes[i]
		switch shape.(type) {
		case *Circle:
			cntCircle++
			b.remove(i)
		default:
			i++
		}
	}

	if cntCircle == 0 {
		return fmt.Errorf("circles doesn`t exist in this slice")
	}
	return nil
}
