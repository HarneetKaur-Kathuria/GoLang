package main

import "fmt"

func main() {

	c := make(chan int)
	gr := 1

	for gr < 11 {

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}

		}()
		gr++

	}

	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}

}
