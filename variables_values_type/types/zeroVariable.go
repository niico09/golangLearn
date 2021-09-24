package main

import (
	"fmt"
)

var y string
var z int

func main() {
	fmt.Println("priting the value of y:", y, ", ending.")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println("int value", z)
	fmt.Printf("%T\n", z)
}
