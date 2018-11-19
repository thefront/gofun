package main

import (
	"fmt"
)

var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true

func main() {
	fmt.Println("Hello, Hands on Exercise 3")

	s := fmt.Sprintf("%x\t%s\t%t", x3, y3, z3)
	fmt.Println(s)

	g := fmt.Sprintf("%v\t%v\t%v", x3, y3, z3)
	fmt.Println(g)

}
