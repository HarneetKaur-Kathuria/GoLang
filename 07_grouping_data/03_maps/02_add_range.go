package main

import "fmt"

func main() {

	m := map[string]int{
		"Hello":  12,
		"Winner": 15,
	}

	fmt.Println(m["Winner"])

	// add pair to map

	m["Singer"] = 29

	// iterate the map

	for k, v := range m {
		fmt.Printf("Key : %v\t with  Value  : %v", k, v)
		fmt.Println("")
	}

}
