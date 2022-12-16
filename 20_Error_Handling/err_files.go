package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("name.txt") // os package to Create Files
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	r := strings.NewReader("Hello I am ")
	io.Copy(f, r) // to write the content to file
}
