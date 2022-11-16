package main

import "fmt"

func main() {

	m := map[string] string{
		"bond": "hello I am here",
		"cute":  "Hi I am so cute",
	}
	fmt.Println(m)


	mp := map[string][]string{
		"bond": [] string {"hello I am here", "hi there"},
		"cute": []  string{"Hi I am so cute", "cuties"},
	}

	fmt.Println(mp)


	for k , v := range mp {
		fmt.Printf("the record fo Key  :  %v\n", k)
			for i , elm := range v {
				fmt.Printf("\t \t Index %v\t element is %v", i , elm)
				fmt.Println("")
			}
	}



}
