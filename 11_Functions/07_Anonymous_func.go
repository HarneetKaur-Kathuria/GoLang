package main

import "fmt"

func main() {
	foo()

	// Anonymous Function aka Self executing function
	func() {
		fmt.Println("Hello I am Anonymous Function")
	}()

	// Anonymous func with argument

	func(x int) {
		fmt.Println("Anonymous Function with arg of int ", x)
	}(15)

}

func foo() {
	fmt.Println("I am foo")
}
