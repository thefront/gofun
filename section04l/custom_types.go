package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	fmt.Println("Hello, playground")
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b)  // converting example, it's called this in go not casting
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
