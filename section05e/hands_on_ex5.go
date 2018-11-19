package main

import (
	"fmt"
)

type own int

var x5 own
var y5 int

func main() {
	fmt.Println("Hello, Hands on Exercise 5")

	fmt.Printf("x is of type %T\n", x5)
	fmt.Println(x5)
	x5 = 42
	fmt.Println(x5)
	y5 = int(x5)
	fmt.Printf("y5 is of type %T\n", y5)
}
