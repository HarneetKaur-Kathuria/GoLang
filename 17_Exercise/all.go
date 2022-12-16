package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
func main() {
	
	wg.Add(2)
	go func (){
		
		for i:=0; i<5; i++{
			runtime.Gosched()
			fmt.Println("Hello : ", i )
			wg.Done()
		}
		// wg.Done()
	}()

	// go greet()
	go func(){
		for i:=0; i<5; i++{

			fmt.Println("Greet : ", i )
		}
		wg.Done()
	}()
	wg.Wait()

}

func greet (){

	for i:=0; i<5; i++{
		fmt.Println("Greet : ", i )
	}
	wg.Done()
}




