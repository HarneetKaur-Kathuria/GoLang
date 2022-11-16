package main

import "fmt"


// all globle var has scope limited to package
var name string = "Winner"
var aim = "Famous"
var string_backquotes = `Hello, "I am the most famous"`

var string_backquotes2 = `Hello, 


"I am the most famous"`
func main() {
	fmt.Println(name)
	fmt.Println(aim)

	fmt.Println(string_backquotes)
	fmt.Println("")
	fmt.Println(string_backquotes2)

	// can not reassign with the different TYPE
	// :ERROR: cannot use 125 (untyped int constant) as string value in assignment
	// name = 125
	// fmt.Println(name)
	

	//Print f
	fmt.Printf("%T     ", name)
	fmt.Printf("%T \n", aim)


}
