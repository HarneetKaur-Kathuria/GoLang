package main

import "fmt"

type own int

var x own

func main() {

	fmt.Println(x)
	x = 125

	fmt.Printf("%T\n", x)
	fmt.Println(x)
}
