package main

import "fmt"

type liking struct {
	first     string
	last      string
	ice_cream []string
}

func main() {
	l1 := liking{
		first:     "Win",
		last:      "winner",
		ice_cream: []string {"vanilla", "chocolate",},
	}

	
	fmt.Println(l1.first, l1.last)
	for i , v := range l1.ice_cream	{
		fmt.Println(i ,v)
	}
}
