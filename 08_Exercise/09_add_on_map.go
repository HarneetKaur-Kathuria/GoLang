package main

import "fmt"

func main() {
	mp := map[string][]string{
		"bond": []string{"hello I am here", "hi there"},
		"cute": []string{"Hi I am so cute", "cuties"},
	}

	fmt.Println(mp)

	mp["famous"] = []string{"I ", "am", "Happy"}

	for k, v := range mp {
		fmt.Printf("the record fo Key  :  %v\n", k)
		for i, elm := range v {
			fmt.Printf("\t \t Index %v\t element is %v", i, elm)
			fmt.Println("")
		}
	}
}
