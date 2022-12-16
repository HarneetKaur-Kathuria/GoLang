package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Welcome to the Goroutine Session")
	wg.Add(2)
	go greet()
	go Hello()
	wg.Wait()
	fmt.Println("Its all About Gorutine")
}

func greet() {
	fmt.Println("I am at Greet")
	wg.Done()
}

func Hello() {
	fmt.Println("I am at Hello")
	wg.Done()
}
