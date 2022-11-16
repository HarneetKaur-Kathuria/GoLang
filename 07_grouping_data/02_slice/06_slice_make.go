package main

import "fmt"

func main() {

	x := make([]int, 5, 10) // make ( type , leng, cap )
	// cap is kind of underlying array with cap 10
	// because slice is build on arrays its kind of arraylist
	// the moment we add 11th element the the cap will get doubled

	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	x = append(x, 25, 45)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	x = append(x, 8, 9, 10)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	// adding 11th element
	x = append(x, 11)
	fmt.Println(x)
	fmt.Println(len(x), cap(x)) // length : 11 , capacity : 20 (10*2)
}
