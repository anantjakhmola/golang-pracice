//Print a value in decimal, binary and hex

package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Printf("%d,%b,%#x", x, x, x)
}
