package shape

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "RightTriangle", shape: RightTriangle{10, 50}, want: 250},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T){
			got := areaTest.shape.Area()
			if got != areaTest.want {
				t.Errorf("%#v got '%.2f' want '%.2f'", areaTest.shape, got, areaTest.want)
			}
		})
	}

	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != want {
	// 		t.Errorf("got '%.2f' want '%.2f'", got, want)
	// 	}
	// }

	// t.Run("rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	checkArea(t, rectangle, 100.0)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}