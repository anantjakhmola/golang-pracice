package main

import (
	"encoding/json"
	"fmt"
)

//if I change the first character of fields to capital the marshal will work correctly
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 :=
		person{
			First: "James",
			Last:  "BOnd",
			Age:   21,
		}

	p2 :=
		person{
			First: "Mary",
			Last:  "mob",
			Age:   30,
		}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
