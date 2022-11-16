package main

import "fmt"

func main() {

	if true {
		fmt.Println("Winner")
	}

	if false {
		fmt.Println("Not Winner")
	}

	if !true {
		fmt.Println("Not Winner")
	}

	if !false {
		fmt.Println("Famous")
	}

	if 2 == 2 {
		fmt.Println("Famous")
	}

	if 3 != 5 {
		fmt.Println("Famous")
	}

	if "winner" == "winner" {
		fmt.Println("Famous")
	}

	if "winner" == "famous" {
		fmt.Println("Winner")
	}

}
