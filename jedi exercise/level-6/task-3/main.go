package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("yoyo")
}

func foo() {
	fmt.Println("hey hey")
}
