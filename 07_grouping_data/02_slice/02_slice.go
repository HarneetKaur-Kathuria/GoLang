package main

import "fmt"

func main() {

	// x:= type {values} // composite literal

	x := []int{4, 5, 6, 7, 8} // range is not necessary
	fmt.Println(x)
	fmt.Println(len(x))

	// SLICE allows to group data of SAME TYPE
	// slice is of dynamic size
}
