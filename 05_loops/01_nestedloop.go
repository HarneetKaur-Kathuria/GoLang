package main

import "fmt"

func main() {
	/*
		for i := 0; i < 5; i++ {

			for j := 0; j < 2; j++ {
				fmt.Printf("The outer loop %d\t  :  The inner loop %d", i, j)
				fmt.Println("")
			}
		}
	*/
	for i := 0; i < 5; i++ {

		fmt.Printf("The outer loop %d", i)
		fmt.Println("")

		for j := 0; j < 2; j++ {
			fmt.Printf("\t \t The inner loop %d", j)
			fmt.Println("")
		}
	}

}
