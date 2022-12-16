package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 24
		close(c)
	}()
	// close(c)
	v, ok := <-c
	// close(c)
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}
