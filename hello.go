package main

import "fmt"

var item string

func main() {
	fmt.Println(item)
}

func init() {
	item = "초기화"
}
