package main

import (
	"fmt"
	"runtime"
)

var n int
var f float64

func main() {

	x := 25
	y := 25.25

	fmt.Printf("%v is of type      	: %T \n", x, x)
	fmt.Printf("%v is of type      : %T \n", y, y)
	fmt.Println("")

	n = 45
	f = 45.45

	fmt.Printf("%v is of type      	: %T \n", n, n)
	fmt.Printf("%v is of type       : %T \n", f, f)

	/*
		n = 45.45 // Error: cannot use 45.45 (untyped float constant) as int value in assignment
		f = 45.45

		fmt.Printf("%v is of type      	: %T \n", n, n)
		fmt.Printf("%v is of type       : %T \n", f, f)
	*/

	fmt.Println("")

	var unit_int uint8 = 250 // unsigned int
	var int_8 int8 = -128    //signed int
	fmt.Printf("%v is of type       : %T \n", unit_int, unit_int)
	fmt.Printf("%v is of type       : %T \n", int_8, int_8)

	// var int_8 int8 = -129  // error : as int8 value in variable declaration (overflows)

	/*------------runtime ------------------*/

	fmt.Println(runtime.GOOS)   // operting System
	fmt.Println(runtime.GOARCH) // 64bits

}
