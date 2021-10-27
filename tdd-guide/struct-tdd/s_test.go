package s

import "testing"

// func TestRectangle(t *testing.T) {
// 	r := Rectangle{10.0, 10.0}
// 	t.Run("get rectangle perimeter", func(t *testing.T) {
// 		got := r.Perimeter()
// 		want := 40.0

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// 	t.Run("get rectangle area", func(t *testing.T) {

// 		got := r.Area()
// 		want := 100.0

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})
// }

// func TestCircle(t *testing.T) {
// 	c := Circle{10.0}
	// t.Run("get circle perimeter", func(t *testing.T) {
	// 	got := c.Perimeter()
	// 	want := 125.66

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

// 	t.Run("get circle area", func(t *testing.T) {

// 		got := c.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 		}
// 	})
// }

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{10, 10}, 40.0},
		{Circle{20}, 125.66370614359172},
		{Triangle{10, 10}, 0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %f want %.2f", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12, 6}, 72},
		{Circle{10}, 314.1592653589793},
		{Triangle{10, 10}, 50.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
