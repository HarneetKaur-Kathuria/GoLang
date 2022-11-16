package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello I am Println")
	fmt.Fprintln(os.Stdout, "Hello I am stdOut")               // download "os package"
	io.WriteString(os.Stdout, "Hello I am writer Interface") // download "io" package
}
