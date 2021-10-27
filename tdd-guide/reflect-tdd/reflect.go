package main

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	// need trans x to string
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field((i))

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func main() {
	t := func(input string) {
		fmt.Println(input)
	}
	walk("a", t)
	walk(1, t)
	walk(1.2, t)
	walk([]string{"hello"}, t)
}
