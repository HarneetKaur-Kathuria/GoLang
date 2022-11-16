package main

import "fmt"

func main() {

	s, sum := total1("Hello ", 2, 5, 5, 6, 8, 9)
	fmt.Println(s, sum)
}

func total1(s string, x... int) (string, int) {
	sum := 0
	fmt.Printf("%T\n",x)
	fmt.Println(len(x))
	fmt.Println(x)

	// v here contains the index value	
	// for  v := range x {
	// 	sum += v
	// 	fmt.Println( v, sum)
	// }

	for i:=0; i<len(x); i++{
		sum+=x[i]
	}
	// for i, v := range x {
	// 	sum += v		
	// 	fmt.Println(i , v)
	// }
	fmt.Println(sum)
	return s, sum
}
