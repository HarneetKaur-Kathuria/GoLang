package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	sum := sum(nums...)
	fmt.Println("Sum is ", sum)

	sum1 := bar(nums)
	fmt.Println("Sum is ", sum1)
}

func sum(x ...int) int {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}

func bar( x[] int) int{
	sum :=0;
	for _,v := range x{
		sum += v
	}
	return sum
}
