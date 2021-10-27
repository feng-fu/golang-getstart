package main

import "fmt"

type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}

type snake struct {
	Type      string
	Poisonous bool
}

func (s snake) description() string {
	return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func (c cat) description() string {
	return fmt.Sprintf("Sound: %v", c.Sound)
}

func main() {
	// var multiD [2][3]int
	// var z = make([]int, 8)
	// z = append(z, 1, 2, 3)
	// fmt.Println("Hello", z, len(z), cap(z))
	// t := 10
	// increment(&t)
	// fmt.Println(t)
	var a animal
	a = snake{Poisonous: true}
	fmt.Println(a.description())
	a = cat{Sound: "Meow!!!"}
	fmt.Println(a.description())
}
