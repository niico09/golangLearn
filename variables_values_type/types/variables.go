package main

import (
	"fmt"
)

var y = 42

//	DECLARE the VARIABLE with the IDENTIFIER "z"
//	is of TYPE string
//	and I ASSIGN the VALUE "Shaken, not sirred"

var z string = "Shaken, not stirred"
var a = `James said, "Shaken, 


not stirred"`

func main() {
	fmt.Print(y)
	fmt.Printf("\n%T\n", y)
	fmt.Print(z)
	fmt.Printf("\n%T\n", z)
	fmt.Print(a)
	fmt.Printf("\n%T\n", a)
}
