package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup // variable wg is of Type WaitGroup From package Sync

	fmt.Println("go routines with Anonymns functions")
	wg.Add(2)
	fmt.Println("begin CPU", runtime.NumCPU())
	fmt.Println("begin GoRoutines", runtime.NumGoroutine())
	go func() {
		fmt.Println("1st GoRoutines")
		wg.Done()
	}()

	go func() {
		fmt.Println("2nd GoRoutines")
		wg.Done()
	}()
	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid GoRoutines", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end GoRoutines", runtime.NumGoroutine())

	fmt.Println("Its done")
}
