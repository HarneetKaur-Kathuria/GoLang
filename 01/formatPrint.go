package main

import "fmt"

var n = 25

func main() {
	fmt.Print("the number is \n", n)
	fmt.Printf("the number in  binary form %b\n", n)
	fmt.Printf("the number in hexaDecimal form %x \n", n)

	// n = 197
	fmt.Printf("the number in hexaDecimal form %#x \n", n)
	fmt.Printf("the number in  binary form %#b\n", n)
	fmt.Printf("the number is %#v\n", n) // %v is for normal value 

	
	fmt.Printf("the number form %#b\t %b \t %x \t %#x \n", n, n, n, n)


	// Sprintf : covert the input into a string

	s := fmt.Sprintf("the number form %#b\t %b \t %x \t %#x \n", n, n, n, n)

	fmt.Print(s)

}
