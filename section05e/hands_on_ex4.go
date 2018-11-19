package main

import (
	"fmt"
)

type own int

var x4 own

func main() {
	fmt.Println("Hello, Hands on Exercise 4")

	fmt.Printf("x is of type %T\n", x4)
	fmt.Println(x4)
	x4 = 42
	fmt.Println(x4)

}
