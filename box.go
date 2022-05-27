package golang_united_school_homework

import (
	"fmt"
)

const shapeNotFoundMessage = "shape wasn't found"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapes:         []Shape{},
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity <= len(b.shapes) {
		return fmt.Errorf("out of range capacity")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if err := b.checkRange(i); err != nil {
		return nil, err
	}
	shape := b.shapes[i]
	if shape == nil {
		return nil, fmt.Errorf(shapeNotFoundMessage)
	}
	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if err := b.checkRange(i); err != nil {
		return nil, err
	}
	shape := b.shapes[i]
	if shape == nil {
		return nil, fmt.Errorf(shapeNotFoundMessage)
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if err := b.checkRange(i); err != nil {
		return nil, err
	}
	sh := b.shapes[i]
	if shape == nil {
		return nil, fmt.Errorf(shapeNotFoundMessage)
	}
	b.shapes[i] = shape
	return sh, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	isCircleExist := false
	for i, shape := range b.shapes {
		switch shape.(type) {
		case Circle:
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			isCircleExist = true
		}
	}
	if !isCircleExist {
		return fmt.Errorf("circle is not exist")
	}
	return nil
}

func (b *box) checkRange(i int) error {
	if i < 0 || i > len(b.shapes)-1 {
		return fmt.Errorf("index %d is out of range of array", i)
	}
	return nil
}
