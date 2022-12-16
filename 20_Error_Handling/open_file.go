package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("name.txt")

	if err != nil {
		fmt.Println("checking  ", err, f)
		return
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs)) // converting byte code to string
}
