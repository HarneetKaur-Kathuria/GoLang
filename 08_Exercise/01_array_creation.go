package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x)
	fmt.Printf("%v\t is of type :\t%T\n", x, x)

	// when siz eis specified
	var y [5]int
	fmt.Println(y)
	fmt.Printf("%v\t is of type :\t%T\n", y, y)

	// another way
	z := [5]int{5, 4, 6, 9, 1}
	for i, v := range z {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", z)
	fmt.Println("")

}
