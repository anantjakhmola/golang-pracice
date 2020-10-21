package main

import "fmt"

type person struct {
	first string
}

func changeMe(p *person) {
	p.first = "MIss moneypenny"
	//(*p).first = "MIss moneypenny"

}

func main() {
	p1 := person{first: "JAmes BOnd"}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}
