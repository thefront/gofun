package main

import (
	"fmt"

	"github.com/thefront/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!niagA ,olleH"))
	foo()
}

func foo() {
	fmt.Println("hello again on a new line")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
