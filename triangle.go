package golang_united_school_homework

import "math"

var squaredThreeDividedFour = math.Sqrt(3) / 4

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (t Triangle) CalcArea() float64 {
	return squaredThreeDividedFour * math.Pow(t.Side, 2)
}
