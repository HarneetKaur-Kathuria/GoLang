package main

import "fmt"

var x int
var g func() = func(){
	fmt.Println("Outside main")
}

var k func()

func main() {

	s :=func() {
		fmt.Println("Hello I am Anonymous")
	}
	s()
	fmt.Printf("%T\n", s)
	fmt.Println(x)

	test :=func() {
		x :=0
		for i:=0; i<10; i++{
			fmt.Print(i, " ")
			x +=i
		}
		fmt.Println("")
		fmt.Println("Total is ", x)
	}
	test()

	k = test // assigining one func to another
	k() // calling with the var name
	fmt.Printf("%T\n", test)
	g()
}
