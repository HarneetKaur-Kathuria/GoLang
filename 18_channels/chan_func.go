package main

import "fmt"

func main() {
	ch := make(chan int)

	// send value
	go foo(ch)

	//recieve val
	bar(ch)
	fmt.Println("Its DONE")
}

// send
func foo(c chan<- int) {
	c <- 25
	c <- 24 
}

func bar(c <-chan int) {
	fmt.Println("chan val", <-c)
	fmt.Println("chan val", <-c)
}
