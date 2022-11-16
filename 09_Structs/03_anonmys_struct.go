package main

import "fmt"

func main() {

	// creating a struct with out name 
	p1 := struct {
		first_name string
		last_name  string
		age int
	}{
		first_name: "Win",
		last_name:  "winner",
		age : 19,
	}

	fmt.Println(p1)
	fmt.Println(p1.first_name , p1.last_name)


}
