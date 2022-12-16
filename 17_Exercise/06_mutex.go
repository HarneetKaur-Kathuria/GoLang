package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Incremeter with GoRoutines")
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			mu.Lock()
			temp := count // read the value and store in new variable
			temp++        // increment the value
			count = temp
			fmt.Print(count , " ")
			mu.Unlock()
			wg.Done()
		}()
	
	}
	fmt.Println("GoRoutines ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Counter", count)
}
