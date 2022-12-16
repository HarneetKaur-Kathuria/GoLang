package main

import "fmt"

func main() {
	fmt.Println("Welcome to the channels buffer")

	ch := make(chan int, 1) // chan can have 2 values now //
	ch <- 25                // but can listen to one , its just it wont throw the deadloack err

	fmt.Println("chan", <-ch)

}
