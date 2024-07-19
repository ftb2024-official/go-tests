package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length, Width float64
}

type Circle struct {
	Raduis float64
}

type Square struct {
	Side float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Raduis * c.Raduis
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}
