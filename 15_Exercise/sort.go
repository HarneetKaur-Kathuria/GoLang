package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{12, 54, 11, 12, 5, 24, 42, 6, 15}
	str := []string{"Hi", "Hello", "Best", "Better", "Cute", "Perfect"}

	fmt.Println("Unsorted : ", x)
	sort.Ints(x)
	fmt.Println("Sorted   : ", x)

	fmt.Println("Unsorted : ", str)
	sort.Strings(str)
	fmt.Println("Sorted   : ", str)

}
