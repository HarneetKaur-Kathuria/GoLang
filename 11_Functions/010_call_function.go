package main

import "fmt"

func main() {
	a := HelloPrint
	b := helloBye
	print(a)
	print(b)
}

func HelloPrint(name string) {
	fmt.Printf("Hello  %+v\n", name)
}

func helloBye(name string) {
	fmt.Printf("Hi %v\n", name)
}

func print(f func(s string)) {
	f("Win")
}
