package main

import "fmt"

func main() {

	g := func(xi []int) int {
		sum := 0
		for i := 0; i < len(xi); i++ {
			sum += xi[i]
		}
		return sum
	}
	res := foo(g, []int{1, 2, 5, 6})
	fmt.Println(res)

}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
