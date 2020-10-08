package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not work")
	case (2 == 4):
		fmt.Println("this should not work also")
	case (4 == 4):
		fmt.Println("this should work")
		fallthrough
	case (5 == 5):
		fmt.Println("this will also work")
	}

}

//Also check this out
//https://play.golang.org/p/De_2ZfSaqN
