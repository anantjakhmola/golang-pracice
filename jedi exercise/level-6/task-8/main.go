package main

import "fmt"

func main() {
	fmt.Println(sum())
	f := foo()
	fmt.Println(f())
}

func sum() int {
	return act()
}

func act() int {
	a := 42
	b := 33
	return a + b
}

//! OR

func foo() func() int {
	return func() int {
		return 43
	}
}
