package shape

type RightTriangle struct {
	Base float64
	Height float64
}

func (r RightTriangle) Area() float64 {
	return (r.Base * r.Height) * 0.5
}