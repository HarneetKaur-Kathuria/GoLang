package main

import "fmt"

func main() {

	a := 2
	fmt.Printf("%d \t %b \t %#U\n", a, a, a)

	b := a << 1 // a*2
	fmt.Printf("%d \t %b \t %#U\n", b, b, b) //%d is for decimaln value


	c := a >> 1 // a/2 
	fmt.Printf("%d \t %b \t %#U\n", c, c, c) 
}
