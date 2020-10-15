package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretagent struct {
	person
	ltk bool
}

func (s secretagent) speak() {
	fmt.Println("Hi I am ", s.person.first, s.last)
}

func main() {

	s1 := secretagent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(s1)
	s1.speak()

}
