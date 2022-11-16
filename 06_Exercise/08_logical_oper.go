package main

import "fmt"

func main() {

	fmt.Printf("True && True   is   : %v\n", true && true)
	fmt.Printf("True && False  is   : %v\n", true && false)
	fmt.Printf("False && False is   : %v\n", false && false)
	fmt.Printf("True || False  is   : %v\n", true || false)
	fmt.Printf("True || True   is   : %v\n", true || true)
	fmt.Printf("False && False is   : %v\n", false || false)
	fmt.Println(!true)
}
