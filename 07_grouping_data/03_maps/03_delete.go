package main

import "fmt"

func main() {

	m := map[string]int{
		"Hello":  12,
		"Winner": 15,
	}

	fmt.Println(m)

	// delete the key
	delete(m, "Hello")
	delete(m, "Question") // wont throw error

	fmt.Println(m)

	if v , ok := m ["Winner"] ; ok{
		fmt.Println("Winner : ", v)
		delete(m, "Winner")
	}

	fmt.Println(m)

}
