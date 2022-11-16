package main

import "fmt"

func main() {
	a := foo()
	fmt.Println("Function foo Returns an int", a)
	s , x :=  bar()
	fmt.Println(s, x)
	// alternate way 
	fmt.Println(bar())
}

func foo() int {
	return 42
}

func bar () (string, int){
	return "Hello I am String with value ", 24
}
