package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (s person) speak() {
	fmt.Println("Hey my name is", s.first, s.last, "and my age is", s.age)
}

func main() {
	s1 := person{
		first: "James",
		last:  "bond",
		age:   21,
	}
	s1.speak()
}
