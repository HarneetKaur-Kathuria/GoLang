package main

import "fmt"

func main() {
	defer foo()
	bar()

}

func foo() {

	defer func(){
		fmt.Println("I am", 25) // lasr
	}()
		fmt.Println("I am in foo") //2
}

func bar() {
	fmt.Println("I am in Bar") // 1
}
