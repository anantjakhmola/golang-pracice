package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[2])
	for i, v := range x {
		fmt.Println(i, v)
	}
}
