package main

import "fmt"

func main() {
	a := (22 == 22)
	b := (44 <= 23)
	c := (22 >= 21)
	d := (32 != 21)
	e := (21 > 34)
	f := (12 < 131)

	fmt.Println(a, b, c, d, e, f)

}
