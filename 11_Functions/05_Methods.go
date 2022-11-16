package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}
// func (r receiver) identifier (parameters) (returns(s)){ code }

func (s secretAgent) show (){
	fmt.Println("Hello I am ", s.first,s.last)
}

func (s person) show (){
	fmt.Println("Hello I am ", s.first,s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Win",
			last:  "Winner",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	// creating method to link 
	sa1.show()

	p1 := person{
		first: "Good",
		last: "Life is so good",
	}

	fmt.Println(p1)
	p1.show()
}
