package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for val := range c {
		fmt.Println(val)
	}
}
