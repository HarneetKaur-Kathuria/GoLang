package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
}

func main() {

	b := Book{
		Title:  "Go Lang",
		Author: "Todd",
	}

	lib := []Book{b}

	fmt.Println(lib) //

	bs, err := json.Marshal(lib) //  encoding data to json : marshalling
	if err != nil {              // returns bytes and err
		fmt.Println(err)
	}

	fmt.Println(string(bs)) // conversion : byte code to string

	// output : [{}]
	// when the struct fields are lower case (means limited scope)
	// struck name could be in any case

}
