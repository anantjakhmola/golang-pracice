package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type agent struct {
	person
	ltk bool
}

func main() {
	sa1 := agent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   21,
		},
		ltk: true,
	}
	p2 := person{
		first: "Jane",
		last:  "Mary",
	}
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.last)

}
