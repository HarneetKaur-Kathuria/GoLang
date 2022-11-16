package main

import "fmt"

func main() {

	s := []int{25, 50, 28, 1, 11, 111}
	fmt.Println(s)

	s = append(s, 8, 7, 6, 5)
	fmt.Println(s)

	slice := []int{5, 6, 7, 8}
	fmt.Println(slice)

	s = append(s, slice...)
	fmt.Println(s)

}
