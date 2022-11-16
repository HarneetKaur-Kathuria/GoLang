package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d kb \t \t  in binary : %b\n", kb, kb)
	fmt.Printf("%d mb \t \t  in binary : %b\n", mb, mb)
	fmt.Printf("%d gb \t \t  in binary : %b\n", gb, gb)
}
