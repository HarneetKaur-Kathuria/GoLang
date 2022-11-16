package main

import "fmt"

func main() {
	x := 0
	fmt.Println("Value of x before	: ", x)
	fmt.Println("Addre of x before	: ", &x)

	foo(&x)
	fmt.Println("Value of x after	: ", x)
	fmt.Println("Addre of x after	: ", &x)
}

func foo(ad *int){
	fmt.Println("Before : ", *ad)
	*ad = 24 //value at address is 24 
	fmt.Println("After  : ", *ad)
}