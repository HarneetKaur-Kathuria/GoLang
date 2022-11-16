package main

import "fmt"

func main() {

	slice := []int{25, 45, 48, 79, 257}

	fmt.Println(slice)
	fmt.Println(slice[:])
	fmt.Println(slice[0:2])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:])

}
