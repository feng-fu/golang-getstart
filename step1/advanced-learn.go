package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type age int

func (a age) sayage() {
	fmt.Println("this year's age is ", a)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "RED", "YELLOW", "RED"}
	return strings[c]
}

func main() {
	// r1 := Rectangle{10, 20}
	// r2 := Rectangle{20, 30}
	// c1 := Circle{14}
	// c2 := Circle{100}
	// fmt.Println("Area of r1 is:", r1.area())
	// fmt.Println("Area of r2 is:", r2.area())
	// fmt.Println("Area of c1 is:", c1.area())
	// fmt.Println("Area of c2 is:", c2.area())

	// var a = age(10)
	// a.sayage()

	// fmt.Println(WHITE, BLACK)

	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We hav %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is ", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")

	boxes.PaintItBlack()

	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is ", boxes.BiggestColor().String())
	fmt.Println("Obviously, now, the biggest one is ", boxes.BiggestColor())

}
