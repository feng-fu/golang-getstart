package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}

	var i []interface{}
	for _, value := range array {
		i = append(i, value)
	}

	fmt.Println(i...)
	// fmt.Println(array...)
}