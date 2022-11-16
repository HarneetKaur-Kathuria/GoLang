package main

import "fmt"


// known as COMPOSIT Data Structure 
type person struct {
	first_name string
	last_name  string
	age int
}

func main() {

	// value of type person
	p1 := person{
		first_name: "Win",
		last_name:  "winner",
		age : 19,
	}

	fmt.Println(p1)
	fmt.Println(p1.first_name , p1.last_name)

	p2 := person {
		first_name: "Fame",  
		last_name: "Famous",
	}

	fmt.Println(p2)
	fmt.Println(p2.first_name , p2.last_name)

}
