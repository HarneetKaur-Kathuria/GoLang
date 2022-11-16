package main

import "fmt"

type person struct {
	name string
	age  int
}
type ByAge [] person
func (a ByAge) len() int {return len(a)}
func (a ByAge) Swap(i, j int) { a[i], a[j]}
func main() {
	p1 := person{name: "Win", age: 21}
	p2 := person{name: "Winner", age: 45}
	p3 := person{name: "Best", age: 24}
	p4 := person{name: "Queen", age: 12}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)

}
