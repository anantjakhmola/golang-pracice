package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := foo(x...)
	fmt.Println(s)

	x2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := bar(x2)
	fmt.Println(s2)
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
