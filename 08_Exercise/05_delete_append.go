package main

import "fmt"

func main() {

	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	s = append(s[2:3], s[6:]...) // only two arguments are allowed
	fmt.Println(s)
}
