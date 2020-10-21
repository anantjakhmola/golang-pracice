package main

import "fmt"

func main() {
	x := 55
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)

}

func foo(y *int) {
	fmt.Println("y before", y)

	fmt.Println("y before", *y)
	*y = 44

	fmt.Println("y after", y)
	fmt.Println("y after", *y)
	*y = 3333
	fmt.Println("Some test", *y)

}
