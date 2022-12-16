package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "Password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pass)
	fmt.Println(bs)


	loginPass := "Password124"

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err !=nil{
		fmt.Println("Can not Log In")
		return
	}

	fmt.Println("Congrats , You are logged in")
	

}
