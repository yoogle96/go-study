package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("grettins: ")
	log.SetFlags(0)

	names := []string{"Yoogle", "Gladys", "Samantha"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
