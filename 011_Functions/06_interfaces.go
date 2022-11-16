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

type human interface {
	speak1()
}

func interface_speak1(h human) {
	fmt.Println("Hello I am inside Human Interface", h)
}
func (s secretAgent) speak1() {
	fmt.Println("Hello I am", s.first, s.last, "---Type : SecretAgent")
}

func (s person) speak1() {
	fmt.Println("Hello I am", s.first, s.last, "---Type : Person")
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "Win",
			last:  "Winner",
		},
		ltk: true,
	}
	// fmt.Println(sa1)
	// creating method to link
	sa1.speak1()

	p1 := person{
		first: "Good",
		last:  "Life is so good",
	}
	p1.speak1()
	// to use interface pass arguments

	interface_speak1(sa1);	
	interface_speak1(p1)

}
