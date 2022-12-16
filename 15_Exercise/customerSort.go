package main

import (
	"fmt"
	"sort"
)

// by age
type ByAge []users

func (u ByAge) Len() int {
	return len(u)
}
func (u ByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u ByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

//byname

type ByFirst []users

func (u ByFirst) Len() int {
	return len(u)
}
func (u ByFirst) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u ByFirst) Less(i, j int) bool {
	return u[i].First < u[j].First
}

type users struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := users{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := users{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := users{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	usersLog := []users{u1, u2, u3}

	// fmt.Println(usersLog)
	// nice way to represent
	for _, user := range usersLog {
		fmt.Println(user.First, user.Last, user.Age)
		for _, val := range user.Sayings {
			fmt.Println("\t\t", val)
		}
	}
	sort.Sort(ByAge(usersLog)) // sort by age
	// fmt.Println(usersLog)

	fmt.Println("---------------Sort By Age---------------\n")
	for _, user := range usersLog {
		fmt.Println(user.First, user.Last, user.Age)
		for _, val := range user.Sayings {
			fmt.Println("\t\t", val)
		}
	}

	sort.Sort(ByFirst(usersLog)) //sort by First Name
	// fmt.Println(usersLog)
	fmt.Println("---------------Sort By Name---------------\n")
	for _, user := range usersLog {
		fmt.Println(user.First, user.Last, user.Age)
		for _, val := range user.Sayings {
			fmt.Println("\t\t", val)
		}
	}

	/// sort sayings also
	fmt.Println("---------------Sort Sayings Also---------------\n")

	for _, user := range usersLog {
		fmt.Println(user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, val := range user.Sayings {
			fmt.Println("\t\t", val)
		}
	}
}
