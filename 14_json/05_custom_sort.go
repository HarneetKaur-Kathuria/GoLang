package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName [] Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name< a[j].Name }



type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p1 := Person{Name: "Win", Age: 21}
	p2 := Person{Name: "Winner", Age: 45}
	p3 := Person{Name: "Best", Age: 24}
	p4 := Person{Name: "Queen", Age: 12}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by Age", people)

	sort.Sort(ByName(people))
	fmt.Println("Sorted by Name", people)


}
