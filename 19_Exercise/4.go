package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}
func receive(c, q chan int) {
	for {
		select {
		case val := <-c:
			fmt.Println("Chan C", val)
		case <-q:
			return

		}
	}
}
func gen(q chan int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		q <- 1
		close(c)
		
	}()

	return c
}
