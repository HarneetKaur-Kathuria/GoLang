package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 80}
	fmt.Println(nums)

	sum := total(nums ...)
	fmt.Println("Total of all numbers are ", sum)


}

func total (n... int)int {
	sum :=0

	for _,v:= range n{
		sum +=v
	}

	return sum
}

fun even ()
