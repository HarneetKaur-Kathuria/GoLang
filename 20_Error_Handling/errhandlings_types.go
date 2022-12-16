package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("names.txt")

	if err != nil {
		// fmt.Println("Error Handling  ", err) //customized err
		// log.Println("Error Handling  ", err) // provides log as well
		// log.Fatalln(err) // don't let the defer func to run and gives : exit status
		log.Panicln(err) //terminates the current goroutine and also the defer function
		// return
	}

	defer fmt.Println("EXIT")

}
