package main

import "fmt"

func main() {

	kb := 1024
	mb := kb * 1024
	gb := mb * 1024

	fmt.Printf("%d kb \t \t  in binary : %b\n", kb, kb)
	fmt.Printf("%d mb \t \t  in binary : %b\n", mb, mb)
	fmt.Printf("%d gb \t \t  in binary : %b\n", gb, gb)

}
