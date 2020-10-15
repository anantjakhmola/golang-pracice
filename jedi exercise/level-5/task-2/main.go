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

	m := map[string]person{
		p.last:  p,
		p1.last: p1,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.fic {
			fmt.Printf(i, val)
		}
	}

}
