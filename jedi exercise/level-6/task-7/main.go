package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("hi function")
	}

	a()
	fmt.Printf("%T\n", a)
}
