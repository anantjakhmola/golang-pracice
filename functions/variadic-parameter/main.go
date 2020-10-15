package main

import "fmt"

func main() {
	foo(2, 3, 4, 6, 3, 21)

}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T", x)
}
