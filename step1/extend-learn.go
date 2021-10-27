package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

type Singer struct {
	Human
	favorite string
}

func (h Human) Sayhi() {
	fmt.Printf("Hi I am %s you can call me on %s \n", h.name, h.phone)
}

func (e Employee) Sayhi() {
	fmt.Printf("Hi, I am %s,working at %s. You can call me on %s \n", e.name, e.company, e.phone)
}

func (s Human) Sing(lyrics []string) {
	fmt.Printf("La la la la... \n")
	for _, v := range lyrics {
		fmt.Printf("%s \n", v)
	}
}

func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

type Men interface {
	Sayhi()
	Sing(lyrics []string)
}

func main() {
	mark := Student{Human{"Mark", 20, "404-2021-1233"}, "MIT"}
	sam := Employee{Human: Human{"Sam", 28, "300-1032-1123"}, company: "Apple Inc"}

	fmt.Println("This is student mark:", mark)
	fmt.Println("This is student sam:", sam)
	// mark.Sayhi()
	// sam.Sayhi()
	// sam.Human.Sayhi()

	// var i Men

	// i = mark
	// fmt.Println(("This is Mark, a student:"))
	// i.Sayhi()
	// i.Sing([]string{"Born ", "to"})
	// var a interface{} // 类似于any类型

	// var i int = 5

	// s := "Hello"

	// a = i
	// fmt.Println(a)
	// a = s
	// fmt.Println(a)

}
