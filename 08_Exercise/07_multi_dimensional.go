package main

import "fmt"

func main() {

	names := []string{"Win", "Winner", "Topper"}

	prof := []string{"Singer", "Actor", "Famous"}

	names_prof := [][]string{names, prof}
	fmt.Println(names_prof)

	for i, v:= range names_prof {

		fmt.Println(i, v)
	}

}
