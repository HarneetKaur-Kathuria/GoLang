package main

import "fmt"


func main() {
	fmt.Println("Assigning func to a Variable")

	f := func() {
		fmt.Println("I have assigned function to variable F")
	}

	f() // calling func f

	f2 := func(x int) {
		fmt.Println("I have assigned function to variable F2 with arg : ", x)
	}

	f2(24)


}
