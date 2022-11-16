package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 9}

	fmt.Println(x)

	// loops over the slice and get VALUE and Index both
	// range

	for i, v := range x {
		fmt.Println(i, "\t", v)
	}

}
