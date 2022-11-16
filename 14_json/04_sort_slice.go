package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{12, 21, 15, 12, 42, 24}
	xs := []string{"Hi", "cool", "life"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println("Sorted : ", xi)

	fmt.Println("-------------------")
	fmt.Println("Unsorted : ", xs)
	sort.Strings(xs)
	fmt.Println("Sorted : ", xs)
}
