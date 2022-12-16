package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// mutex solves the race conditions

	counter := 0

	const num = 15
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			// temp := counter
			mu.Lock()
			runtime.Gosched() // allows other go routines to run // counter : 1
			// temp++
			// counter = temp
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)

}
