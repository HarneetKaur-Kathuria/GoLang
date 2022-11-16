package main

import "fmt"

func main() {

	// xi := [] int {2, 5, 6, 12, 25}
	// sum := total (xi...) // can also pass a slice 
	// sum := total(2, 5, 6, 12, 25)
	sum := total () // can also pass it with zero arguments.. 
	// variadic parameters are like 0 or more
	fmt.Println("Total is ", sum) 



}

func total(x ...int) int {

	sum := 0

	for i, v := range x {

		sum += v
		fmt.Println("for item of position ", i, "the value of ", v, "is being added and sum is ", sum)
	}
	return sum

}
