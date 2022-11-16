package main

import "fmt"

func main() {

	// pointers stores value in memory
	a := 25
	// "&a " tells the address
	fmt.Printf("Value : %v , of Type : %T , with address  : %v\n", a, a, &a)

	// Type of an Address
	fmt.Printf("Type of address :  %T\n", &a) // *int : "*" pointer pointed to int

	// we can't assign the pointer to int but to pointer

	// var b int = &a // error cant assign pointer to int
	var b *int = &a 
	c := &a
	fmt.Println(b , "\t", c)

	// to get value from address : derefrence the address

	fmt.Println("Get Value from address" , *c)

	// fmt.Println(*&a) // combining both

	// & gives the address
	// * gives the value stored at value

	*b = 43 // b store the add of value a and here reassigning the value to 
	fmt.Println(a) // after dereferencing it to value , so value assigning to a also changes 

}
