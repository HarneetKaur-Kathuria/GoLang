package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64 //  why FLOAT64 Is mentioned
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {

	s1 := square{
		side: 12,
	}

	// fmt.Println(s1.area())
	info(s1)

	c1 := circle{
		radius: 25,
	}

	// fmt.Println(c1.area())
	info ( c1 )

}
