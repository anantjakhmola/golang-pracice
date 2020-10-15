package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 33
	fmt.Println("Hey the value is", x, "and length are ", len(x))
	//major difference between Printf and Println are about formating
	fmt.Printf("Hey the value is %d and length are %d ", x, len(x))

}
