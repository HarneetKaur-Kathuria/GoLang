package main

import "fmt"

const (
	a = iota
	b
	c
)

const (  // only with "const" it get reset, iota is predefined IDENTIFIER
	d = iota
	e
	f = iota
)
func main() { 

	const (
		g = iota
		h
		i
	)

	fmt.Println(a, "\t", b, "\t", c)
	fmt.Println(d, "\t", e, "\t", f)
	fmt.Println(g, "\t", h, "\t", i)
}
