package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

// CalcPerimeter returns calculation result of perimeter
func (r *Rectangle) CalcPerimeter() float64 {
	return r.Height*2 + r.Weight*2
}

// CalcArea returns calculation result of area
func (r *Rectangle) CalcArea() float64 {
	return r.Height * r.Weight
}
