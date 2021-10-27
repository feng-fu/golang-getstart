package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// for i := 1;i < len(os.Args);i++ {
		// arg := os.Args[i]
	for i, arg:= range os.Args[1:] {
		fmt.Println(i, float64(i))
		s += sep + arg
		sep = ","
	}
	fmt.Println(s);
}