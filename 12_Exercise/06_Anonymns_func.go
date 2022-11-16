package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello I am Anonymous")
	}()

	func() {
		x :=0
		for i:=0; i<10; i++{
			fmt.Print(i, " ")
			x +=i
		}
		fmt.Println("")
		fmt.Println("Total is ", x)
	}()
}
