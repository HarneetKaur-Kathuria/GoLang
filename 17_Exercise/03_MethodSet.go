package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("GoRoutine With Increment Examples")
	var counter = 0
	var wg sync.WaitGroup
	wg.Add(5)
	for i:=0 ; i< 5; i++{
		go func(){
			temp := counter
			temp ++
			counter=temp
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter", counter)

}
