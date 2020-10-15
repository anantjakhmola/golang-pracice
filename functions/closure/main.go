package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Printf("%#x\n", a())
	fmt.Printf("%#x", b())

	//!! Both the functions have different block of addresses
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
