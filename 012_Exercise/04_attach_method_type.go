package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last , "and I am ", p.age, "years old.")
}
func main() {

	p1 := person{
		first: "Win",
		last: "Winner",
		age: 369,
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
	p1.speak() 
	
}