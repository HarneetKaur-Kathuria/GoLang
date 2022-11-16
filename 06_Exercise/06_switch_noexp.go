package main

import "fmt"

func main() {

	// no expressions means true
	switch {
	case true:
		fmt.Println("True")
		// fallthrough
	case false:
		fmt.Println("False")
		// fallthrough : prints the next statement also whether its true or not
	default:
		fmt.Println("No Valid Value")
	}

}
