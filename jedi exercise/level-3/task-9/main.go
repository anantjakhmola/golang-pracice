//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "Tennis"
	switch favSport {
	case "Bowling":
		fmt.Println("Go to Arena")
	case "Skating":
		fmt.Println("Go to park")
	default:
		fmt.Println("No match found")
	}

}
