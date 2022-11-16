package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) changeMe() {

	*&p.first = "F"
	*&p.last = "A"
	*&p.age = 25

}
func main() {

	p1 := person{
		first: "Win",
		last:  "Winner",
		age:   24,
	}
	fmt.Println("Before : ", p1)
	p1.changeMe() // func attached 
	fmt.Println("After  : ", p1)

	// by creating function and passing add as args
	changeMe(&p1)
	fmt.Println("After  : ", p1)
}

func changeMe(p *person){
	p.first = "hello"
	p.last = "Bye"
	p.age = 42
}
