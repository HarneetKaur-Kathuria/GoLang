package main

import "fmt"

func main() {
	function := foo()
	fmt.Printf("%T\n", function)

	// value := function()
	// fmt.Println("Value returns from function ", value)

	fmt.Println(" Method ", function())

	fmt.Println(" Direct Method ", foo()())


}

func foo() func() int {

	return func() int {
		return 24
	}
}
