package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	{
		x := 22
		fmt.Println(x)
		x++
		fmt.Println(x)

	}
	fmt.Println(x)

}

func foo() int {
	x := 5
	x++
	return x

}
