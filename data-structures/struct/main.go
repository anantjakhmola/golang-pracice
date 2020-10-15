package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Jane",
		last:  "Mary",
	}
	fmt.Println(p1.first)
	fmt.Println(p2.last)

}
