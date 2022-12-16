package main

import "fmt"

func main() {
	ch := make(chan int)

	// send val
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}

		close(ch) // closing chan to avoid deadlock err ,  // not race condition 
	}()

	// receive val

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("ITS DONE")
}
