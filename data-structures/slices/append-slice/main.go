package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(x)
	//fmt.Println(x[1:3])
	x = append(x[:2], 5, 6, 7)
	fmt.Println(x)

	y := []int{11, 22, 33, 44}
	x = append(x, y...)
	fmt.Println(x)
}
