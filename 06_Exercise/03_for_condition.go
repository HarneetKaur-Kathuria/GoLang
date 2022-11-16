package main

import "fmt"

func main() {

	x := 2000
	count := 0
	for x < 2023 {
		fmt.Print(x, "\t")
		x++
		count++
	}
	fmt.Println("")
	fmt.Println(count)

	y := 2000

	for {
		if y > 2022 {
			break
		}
		fmt.Print(y, "\t")
		y++
	}

}
