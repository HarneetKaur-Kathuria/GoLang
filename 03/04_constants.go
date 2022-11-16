package main

import "fmt"
/*
const a = 25
const b = 25.25
const c = "Winner"
*/

const (

	 a = 25
	 b = 25.25
	 c = "Winner"
)
func main() {

	fmt.Printf("%v is of type  :  %T\n", a, a)
	fmt.Printf("%v is of type  :  %T\n", b, b)
	fmt.Printf("%v is of type  :  %T\n", c, c)

}
