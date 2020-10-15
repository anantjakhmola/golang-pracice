package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello")
	}
	f()

	f1 := func(x int) {
		fmt.Println("hello", x)
	}
	f1(22)
}
