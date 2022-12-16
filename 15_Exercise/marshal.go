package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name   string `json:"Course Name"`
	Auther string `json:"Writer"`
	Rate   float64 `json:"Price,omitempty"`
}

func main() {

	b1 := Book{
		Name:   "Go lang",
		Auther: "Kim",
		
	}

	b2 := Book{
		Name:   "Java",
		Auther: "Robert",
		Rate:   1525.25,
	}

	lib := []Book{b1, b2}

	fmt.Println("Library ", lib)

	// bs, err := json.Marshal(lib)
	jsaonByte, err := json.MarshalIndent(lib, "", "\t")
	if err != nil {
		panic(err)
	}

	// fmt.Println("Marshal", string(jsaonByte))
	fmt.Printf("%s\n", jsaonByte)

}
