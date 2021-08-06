package shape

type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}