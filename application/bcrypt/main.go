package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginpass := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpass))
	if err != nil {
		fmt.Println("you cant login")
		return
	}
	fmt.Println("cool")
}
