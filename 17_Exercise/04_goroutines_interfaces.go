package main

import "fmt"

type Person struct {
	name string
}

type human interface {
	speak()
}

func (p *Person) speak() {
	fmt.Println("Hello I am ", p.name) // requires pointer
}

func saySomething(h human) {
	h.speak() //1 : calls the speak method from interface 
}

func main() {
	p1 := Person{
		name: "Win",
	}
	// saySomething(p1)// err : speak method has pointer receiver
	saySomething(&p1)
	p1.speak() // person can directly calls the method also
}
