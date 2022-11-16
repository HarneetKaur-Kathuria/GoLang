package main

import "fmt"

// assign values using operators

func main() {

	a := (25 == 25)
	b := (25 != 25)
	c := (25 >= 50)
	d := (25 <= 50)
	e := 25 > 53
	f := 53 > 25

	fmt.Println(a, b, c, d, e, f)

}
