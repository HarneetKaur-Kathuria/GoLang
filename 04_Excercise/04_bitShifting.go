package main

import "fmt"

func main() {

	x := 2
	fmt.Print(x, "\t")
	fmt.Printf("in Decimal : %d\t| in Binary : %b\t| in Hexa Decimal : %#x\n", x, x, x)

	y := 2 << 1 // x*2
	fmt.Print(y, "\t")
	fmt.Printf("in Decimal : %d\t| in Binary : %b\t| in Hexa Decimal : %#x\n", y, y, y)

	z := 2 >> 1 // y/2
	fmt.Print(z, "\t")
	fmt.Printf("in Decimal : %d\t| in Binary : %b\t| in Hexa Decimal : %#x\n", z, z, z)

}
