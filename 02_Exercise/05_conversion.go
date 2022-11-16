package main

import "fmt"

type own int

var x own
var y int

func main() {

	fmt.Println(x)
	x = 125
	fmt.Println(x)
	fmt.Printf( "%v  type : %T \n", x, x) // main.own

	y = 200;
	fmt.Println(y)
	fmt.Printf("%v type : %T \n", y, y)
	

			/*------------CONVERSION --------------*/

	/* -------- own type to underling type--------------*/

	y = int (x)
	fmt.Println(y)
	fmt.Printf("%v type : %T \n", y, y)

	// assigning the value to new Var

	z := int (x)
	fmt.Println(z)
	fmt.Printf("%v type : %T \n", z, z)

}