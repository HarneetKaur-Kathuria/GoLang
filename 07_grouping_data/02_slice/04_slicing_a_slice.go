package main

import "fmt"

func main() {

	x := [ ]int{25, 26, 27, 28, 29}

	fmt.Println(x)
	fmt.Println(x [:]) // : iterate at every index and position
	fmt.Println(x[1:]) // prints x[1] and onwards 
	fmt.Println(x[1:3]) // from x[1]  - x[2] onwards
	fmt.Println(x[:3]) // prints till x[2]
}
