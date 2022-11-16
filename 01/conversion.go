package main

import "fmt"

var a int

type hotdog int

var b hotdog // we create a new type with underline type
func main() {

	type hot string //
	var s hot

	fmt.Printf("a is of  : %T \n", a)
	fmt.Printf("b is of  : %T \n", b)
	fmt.Printf("s is of  : %T \n", s)

	// cant bcz both are of diff type

	//a=b // error :cannot use b (variable of type hotdog) as type int in assignment

	// ----- conversion

	a = int(b) // now we convertting bcz the underline type caste of hotdog is int
	fmt.Printf("a is of  : %T \n", a) // now b is converted into a

	

}
