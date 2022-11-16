package main

import "fmt"

func main() {

	names := make([]string, 5, 7)

	fmt.Println(names, "\t", len(names), "\t", cap(names))

	// this way : it changes the cap and adds elements after the leng 5
	// names = append(names, "I", "am", "winner")
	// fmt.Println(names, "\t", len(names), "\t", cap(names))

	name_new := []string{"I", "am", "winner"}

	// puting values in initial slice
	for i, v := range name_new {
		names[i] = v
	}

	// printing the slice but leng remains same
	fmt.Println(names, "\t", len(names), "\t", cap(names))

}
