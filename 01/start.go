package main

import "fmt"

func main() {
	fmt.Println("getting Started with go")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
}

func foo() {
	fmt.Println("done")
}
