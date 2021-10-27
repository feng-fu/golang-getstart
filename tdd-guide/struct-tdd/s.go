package s

import "math"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Triangle struct {
	base float64
	height float64 
}

func (t Triangle) Perimeter() float64 {
	return 0
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}
