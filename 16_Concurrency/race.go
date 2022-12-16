package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// race condition : when two goroutines 

	counter := 0
	const num = 15
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			temp := counter
			runtime.Gosched() // allows other go routines to run // counter : 1
			temp++
			counter = temp
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
	//with out runtime.Gosched : counter : 15
	//with runtime.Gosched : counter : 1
}
