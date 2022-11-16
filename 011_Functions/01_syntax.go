package main

import "fmt"

func main() {
	foo()                     // void
	test("Test")              // call by argument with no return
	s := woo("returns Value") // takes argument and VALUE
	fmt.Println(s)

	x, y, z := doubleReturn("String", 24) // multple arguments and multiple return Values
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

func foo() {
	fmt.Println("Hello I am Foo")
}

func test(s string) {
	fmt.Println("Hello I am ", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello I ", s)
}

func doubleReturn(s string, n int) (string, bool, int) {
	a := fmt.Sprint("Hello I am Returning ", s, " VAlue")
	b := true
	return a, b, n
}
