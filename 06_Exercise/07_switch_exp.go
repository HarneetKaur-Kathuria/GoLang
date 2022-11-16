package main

import "fmt"

func main() {

	favSport := "Badminton"
	switch favSport {
		case "skiing":
			fmt.Println("favourite sport is Skiing")
		case "Tennins":
			fmt.Println("favourite sport is Tennins")
		case "Badminton":
			fmt.Println("favourite sport is Badminton")
		default:
			fmt.Println("No Favorite sport")
	}
}
