package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	//! The output of below statement will include * which is a type
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	//! here * is a operator
	//! this  *  operator will derefernce the address and provide the value associated with that address
	fmt.Println(*b)
	//? & will give the ouput and * will derefernce the output

	fmt.Println(&a)
	fmt.Println(*&a)

}
