package main

import "fmt"

func main() {

	x := "Hello"

	if x == "HellO" {
		fmt.Println("Hello , There")
	} else if x == "Bye" {
		fmt.Println("Bye")
	} else {
		fmt.Println("DONE")
	}

}
