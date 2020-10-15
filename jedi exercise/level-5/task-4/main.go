package main

import "fmt"

// type car struct {
// 	door  string
// 	light bool
// }

func main() {

	p1 := struct {
		door  string
		light bool
	}{
		door:  "5",
		light: true,
	}

	fmt.Println(p1)

}
