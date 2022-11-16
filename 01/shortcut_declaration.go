package main

import "fmt"

// declare var outside body
var globle = 111

// DECLARE there is VARIABLE with the IDENTIFIER "zz"
// and that the VARIABLE with the IDENTIFIER is of TYPE int
// and also assigns the ZERO VALUE to "zz" : which is the default value

var zz int

func main() {
	/*Short declaration Operator
	declare & assign = initilization
	*/
	x := 72
	fmt.Println(x)

	x = 35
	fmt.Println(x, "winner")

	// declare variable using Var
	var z = 25
	fmt.Println(z)

	// statement : using operator
	y := 100 + 25
	fmt.Println(y)

	name := "Winner"
	fmt.Println(name)

	// printing globle variable
	fmt.Println(globle)

	// print the default value of var int
	fmt.Println(zz)

	// func to access the globle var outside the body
	foo()

}

func foo() {
	fmt.Println("again : ", globle, "  ", globle)
}
