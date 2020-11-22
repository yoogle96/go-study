package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	if true {
		defer fmt.Println("TEST")
		fmt.Println("QWER")
	}
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
}
