package main

import "fmt"

func main() {

	if x := 42; x == 2 {
		fmt.Println("Great")
	}

	fmt.Println("Everything is Great")

	if x := 42; x == 42 {
		fmt.Println("Great")
	}

	// fmt.Println(x) error : undefined: x

	fmt.Println("Everything is Great")
}
