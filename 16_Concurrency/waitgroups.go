package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // dec a var of type waitGroup
func main() {
	fmt.Println("WaitGroups")
	wg.Add(2) // add the no of Goroutines we have

	go foo()
	// runtime.Gosched()
	defer wg.Wait() // dont let the main method end
	// fmt.Println("Hello")
	go greet()

	// runtime.Gosched()
	// wg.Wait()
}

func foo() {
	fmt.Println("I am foo")
	wg.Done()
}

func greet() {
	fmt.Println("I am greet")
	wg.Done()
}
