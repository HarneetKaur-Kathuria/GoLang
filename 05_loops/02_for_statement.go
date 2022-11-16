package main

import "fmt"

func main() {

	x := 1
	// for statement with single condition
	// this works like a while loop

	for x < 10 {
		fmt.Print(x, "\t")
		x++
	}
	fmt.Println("")

	// for statement with a clause

	y := 1

	for {
		if y > 9 {
			break
		}
		fmt.Print(y, "\t")
		y++
	}
	fmt.Println("")
	fmt.Println("DONE")

}

/*

1.
for i := 0 ; i<10 ; i++{
	i++
}

2.
a :=0
for a<10{
	a++
}

3.

y := 1

	for {
		if(y >9){
			break
		}
		fmt.Print(y, "\t")
		y++
	}


*/
