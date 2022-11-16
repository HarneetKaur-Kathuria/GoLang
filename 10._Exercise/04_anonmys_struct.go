package main

import "fmt"

func main() {

	t1 := struct {
		vehicle   string
		doors     int
		color     string
		fourWheel bool
	}{
		vehicle:   "Car",
		doors:     25,
		color:     "Blue",
		fourWheel: true,
	}

	fmt.Println(t1)

	// struct of map and slice

	t := struct {
		first   string
		friends    map [string] int
		color [] string
		
	}{
		first  : "Win",
		friends : map [string] int {
			"Heena" : 25,
			"Flower" : 52,
		},
		color : []string {
			"Queen" , 
			"princes",
		},
	}

	fmt.Println(t.first)
	fmt.Println(t.friends)
	fmt.Println(t.color)


}
