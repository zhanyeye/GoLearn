package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    got := Perimeter(rectangle)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    // We are declaring a slice of structs by using []struct with two fields, the shape and the want.
    // Then we fill the slice with cases.
    areaTests := []struct {
        name string
        shape Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
        {name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
    }

    for _, tt := range areaTests {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
            }
        })
    }
}


/**

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			// As you can see, the 'f' has been replaced by 'g', using 'f' it could be difficult to know
			// the exact decimal number, with 'g' we get a complete decimal number in the error message
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

*/