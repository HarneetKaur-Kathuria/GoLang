package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Incremeter with GoRoutines")
	var wg sync.WaitGroup
	count := 0
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			// runtime.Gosched()
			temp := count // read the value and store in new variable
			runtime.Gosched()
			temp++        // increment the value
			count = temp
			fmt.Print(count , " ")
			wg.Done()
		}()
	
	}
	fmt.Println("GoRoutines ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Counter", count)
}
