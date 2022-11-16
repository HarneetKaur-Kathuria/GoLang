package main

import "fmt"

var a int = 124
var s string = "Winner"
var check bool = false

func main() {

	fmt.Println(a) // 0
	fmt.Println(s) // nil : empty
	fmt.Println(check) // false

	str := fmt.Sprintf("Values are  : %v \t %v \t %v", a, s, check)

	fmt.Println(str)

}
