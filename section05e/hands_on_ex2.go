package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println("Hello, Hands on Exercise 2")

	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)
	fmt.Println("these are of zero value")

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
