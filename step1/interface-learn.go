package main

import (
	"fmt"
	"reflect"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("(name: %s - age: %d years)\n", p.name, p.age)
}

func main() {
	// list := make(List, 3)
	// list[0] = 1
	// list[1] = "Hello"
	// list[2] = Person{"Dennis", 20}

	// for index, element := range list {
	// 	// if value, ok := element.(int); ok {
	// 	// 	fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	// 	// } else if value, ok := element.(string); ok {
	// 	// 	fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
	// 	// } else {
	// 	// 	fmt.Printf("list[%d] is of a different type, and its value is %v\n", index, element)
	// 	// }

	// 	switch value := element.(type) {
	// 	case int:
	// 		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	// 	case string:
	// 		fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
	// 	default:
	// 		fmt.Printf("list[%d] is of a different type, and its value is %v\n", index, element)
	// 	}
	// }

	var x float64 = 3.4

	p := reflect.ValueOf(&x)
	v := p.Elem()

	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	v.SetFloat(11)
	fmt.Println("value:", v.Float())
}
