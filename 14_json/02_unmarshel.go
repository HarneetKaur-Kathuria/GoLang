package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

func main() {
	s := `[{"Title":"Go Lang","Author":"Todd"}, {"Title":"Winner","Author":"Win"} ]`
	// string of object book
	bs := []byte(s) //  conversion to slice of byte
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	liberary := []book{}
	// var liberary []book // alternative way and effective way
	fmt.Printf("%T\n", liberary)

	err := json.Unmarshal(bs, &liberary)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All books are : ", liberary)

	for i, v := range liberary {
		fmt.Println("Book Details ", i+1)
		fmt.Println("\t ", v.Title, v.Author)
	}

}
