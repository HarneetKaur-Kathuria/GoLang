package main

import "fmt"

func main() {
	foo()
	defer bar() // defer keyword : function run at the last
	sum()
	cool()
}

func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("I am in Bar")
}

func sum() {
	fmt.Println("I am in Sum")
}

func cool(){
	fmt.Println("I am cool")
}