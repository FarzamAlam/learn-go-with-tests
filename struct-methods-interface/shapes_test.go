package structmethods

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f while want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangle : ", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		got := rectangle.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f while want %.2f", got, want)
		}
	})
	t.Run("circlr : ", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g while want %g ", got, want)
		}
	})
}
