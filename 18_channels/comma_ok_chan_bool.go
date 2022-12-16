package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send

	go send(even, odd, quit)
	//receive

	receive(even, odd, quit)
}

func receive(e, o <-chan int, q <-chan bool) {
	// i, ok := <-q
	// if ok {
	// 	fmt.Println("from comma OK", i, ok)
	// } else {
	// 	fmt.Println("from comma OK", i)
	// }
	
	for {
		select {
		case v := <-e:
			fmt.Println("EVEN", v)
		case v := <-o:
			fmt.Println("ODD", v)

		case i, ok := <-q:
			if ok {
				fmt.Println("from comma OK", i, ok)
			} else {
				fmt.Println("from comma OK", ok)
			}
			return 
		}

	}

}
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	
	// close(q)
}
