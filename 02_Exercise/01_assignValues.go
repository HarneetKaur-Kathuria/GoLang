package main

import "fmt"

func main() {
	x := 42
	y := "Thank You"
	z := true

	fmt.Printf("%v is of type : %T \n", x, x)
	fmt.Printf("%v is of type : %T \n", y, y)
	fmt.Printf("%v is of type : %T \n", z, z)

	// single print statements

	fmt.Print(x, "\t", y, "\t", z, "\n")

	// multiple print statements

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
