package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("grettins: ")
	log.SetFlags(0)

	message, err := greetings.Hello("yoogle")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
