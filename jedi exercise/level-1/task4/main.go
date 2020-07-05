package main

import "fmt"

type anant int

var x anant

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 45
	fmt.Println(x)
}
