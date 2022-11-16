package main

import "fmt"

func main() {
	s := "I am the Winning"

	fmt.Println(s)
	fmt.Printf("%T \n", s)
	// converting String to slice of Bytes

	bs := [] byte (s) // cant write link : byte[] (s) // syntax error

	

	
	fmt.Println(bs)
	fmt.Printf("%v \t %T \n", s, bs)

	str := string (bs)
	fmt.Println(str)
}
