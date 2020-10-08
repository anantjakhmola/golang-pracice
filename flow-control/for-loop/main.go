package main

import "fmt"

func main() { // for init ; condition ; post{}
	for i := 0; i <= 5; i++ {
		fmt.Printf("The outer loop %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tThe Inner Loop: %d\n", j)
		}
	}
} //Print the character from 0 - 200 //https://play.golang.org/p/oDVYaDZVI44
