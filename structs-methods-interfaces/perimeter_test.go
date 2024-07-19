package structsmethodsinterfaces

import "testing"

func TestArea(t *testing.T) {
	shapesTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10, 10}, 100},
		{Circle{10}, 314.16},
		{Square{30}, 900},
	}

	for _, test := range shapesTest {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("got %.2f, want %.2f", got, test.want)
		}
	}
}
