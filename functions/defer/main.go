package main

import "fmt"

func main() {
	defer foo()

	bar()
	foo()

	bar()
}

func foo() {
	fmt.Println("foo")

}

func bar() {
	fmt.Println("bar")
}

//deferred function will run after it exits the block of code
