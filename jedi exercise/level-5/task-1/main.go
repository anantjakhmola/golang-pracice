package main

import "fmt"

type person struct {
	first string
	last  string
	fic   []string
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		fic:   []string{"Chocolates", "Vanilla"},
	}
	p1 := person{
		first: "Finity",
		last:  "Rocker",
		fic:   []string{"Kaju", "Mango"},
	}

	for i, v := range p.fic {
		fmt.Println(i, v)
	}
	fmt.Println(p1.first)

}
