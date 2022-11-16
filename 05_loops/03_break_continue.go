package main

import "fmt"

func main() {

	x := 0

	for {
		if x > 50 {
			break
		}

		if x%2 == 0 {
			fmt.Print(x, "\t")
		}
		x++
	}

	fmt.Println("")

	// for i :=0 ; i<51; i++{
	// 	if i%2 ==0 {
	// 		fmt.Print(i , "\t")
	// 	}
	// }

	for i := 0; i < 51; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Print(i, "\t")
	}
}
