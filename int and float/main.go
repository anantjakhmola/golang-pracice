package main

import "fmt"

var x int32
var y float32

func main() {
	x = 100
	y = 2.121212133
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Println(y)
}

//NOTE:all numeric type are distinct except
// NOTE: BYTE which is alias for unin8
// NOTE:rune which is alias for int32
