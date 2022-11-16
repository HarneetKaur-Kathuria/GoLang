package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)

	x = append(x, 25, 45) // append is a function
	fmt.Println(x)

	/*
		x = append (25, 65, x) // and first argument must be a slice
		fmt.Println(x)

	*/

	y := []int{25, 45}
	y = append(y, x...)
	//y = append (x..., y)  // error : syntax error: unexpected y, expecting )
	fmt.Println(y , " length ", len(y))


	// delete a slice 

	y = append (y[:2], y[5:]...) // index 0--1 && 5.... and rest is deleted
	fmt.Println(y, " length ", len(y))

}
