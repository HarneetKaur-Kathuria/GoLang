package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS) // tell the OS
	fmt.Println("ARCH\t\t", runtime.GOARCH) // tell the 
	fmt.Println("CPU\t\t", runtime.NumCPU()) // no of CPU
	fmt.Println("GoRoutines\t", runtime.NumGoroutine()) // num of Gouroutines

	fmt.Println("")

	go foo()

	time.Sleep(10* time.Microsecond) // waits 
	fmt.Println("Testing Goroutine")

}

func foo(){
	
	fmt.Println("I am in foo")
}
