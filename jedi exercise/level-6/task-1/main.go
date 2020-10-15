package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	m, v := bar()
	fmt.Println(m, v)
}

func foo() int {
	var x int
	x = 12
	return x
}

func bar() (int, string) {
	x1 := 22
	x2 := "Hi"
	return x1, x2
}
