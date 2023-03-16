package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assertCalculation(t, got, want)

}

func TestArea(t *testing.T) {

	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}
	//
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{12.0, 6.0}
	//	//got := rectangle.Area()
	//	want := 72.0
	//	checkArea(t, rectangle, want)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	//got := circle.Area()
	//	want := 314.1592653589793
	//	checkArea(t, circle, want)
	//})

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.name, got, tt.hasArea)
		}
	}
}

func assertCalculation(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
