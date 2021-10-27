package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)


func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	messages, err := greetings.Hellos([]string{"Jack", "Lucy", "Lily"})
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}