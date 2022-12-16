package main

import "fmt"

func main() {
	fmt.Println(sum(2, 5))
	fmt.Println(sum(8, 5))

}

func sum(xi ...int) int { // variadic parameter
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
