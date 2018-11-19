package main

import (
	"fmt"
)

var y = 42

var z string = "this is a string declared"

var a string = `
this is like a block and print all in there
in
with returns and "qutoes"
`
var d = `:=`

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y)
	fmt.Println(z)
	fmt.Printf("z is type of %T\n", z)
	fmt.Printf("z is tyoe of %T\n", a)
	fmt.Println(a)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	fmt.Printf("this is called the short declaration operator %s\n", d)
	fmt.Println(d)
}
