package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("yoogle")
	fmt.Println(message)
}
