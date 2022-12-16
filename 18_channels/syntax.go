package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Channels")

	ch := make(chan int)

	// provide error : bcz listing should be first
	// ch <- 25 // puting value to channel
	// fmt.Println(<-ch)

	go func() {
		fmt.Println("chan", <- ch)
	}()
	// no error : bcz listening func is already there
	ch <- 25 // puting value to channel

}
