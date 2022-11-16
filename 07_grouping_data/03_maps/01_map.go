package main

import "fmt"

func main() {

	m := map[string]int{ // [key_type] value_type
		"James":  32,
		"Honey":  64,
		"winner": 65,
	}

	fmt.Println(m)

	fmt.Println(m["James"]) // giving key : output - value
	fmt.Println(m["Hello"]) // 0 : not existed

	// to check whether the key exits or not
	v, ok := m["hello"] // okay stores boolean
	fmt.Println(v , "\t" , ok)

	// can check with if _condition also

	if v, ok := m["James"]; ok{ // v saves the values && ok true of false
		fmt.Println("Key Exists in Map , with Value : ", v)
	}


}
