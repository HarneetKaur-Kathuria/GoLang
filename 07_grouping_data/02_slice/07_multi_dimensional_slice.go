package main

import (
	"fmt"
)

func main() {

	
	names := [] string {"Winner", "Famous", "Unstoppable"}
	fmt.Println(names)
	fmt.Println(len(names))

	traits := [] string {"Singer", "Famous", "Popular"}
	fmt.Println(traits)
	fmt.Println(len(traits))

	combination := [][] string {names, traits} // slice of slice of strings
	fmt.Println(combination)


}