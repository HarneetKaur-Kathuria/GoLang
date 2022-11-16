package main

import "fmt"

func main() {
	s := func(){
		fmt.Println("Hello there I am here")
	}

	s() // s kindo of become function
	// returning function
	x := bar()
	// fmt.Println(x)
	fmt.Printf("%T\n",x )
	xi := x()
	fmt.Printf("%T\n", xi)
	fmt.Println(xi) 
	fmt.Println(x()) // bar() > x()
	fmt.Println(bar()())// alternate and one line statement

}

	// returning func

	func bar() func() int{
		return func() int {
			return 15
		}
	}
