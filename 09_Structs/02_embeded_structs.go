package main

import "fmt"

type student struct {
	first string
	last  string
	age   int
}

type persons struct {
	student
	present bool
}

func main() {

	p1 := persons{
		student: student{
			first: "hello",
			last:  "There",
			age:   6,
		},
		present: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.student, p1.first, p1.present, p1.age)

}
