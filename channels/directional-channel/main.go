//! CHANNEL CAN BE DIRECTIONAL AND BI-DIRECTIONAL ALSO
package main

import "fmt"

func main() {
	//? c := make(chan int, 2)
	//? This is a send only channel
	//c := make(chan<- int, 2)

	//? Receive only channel
	//c := make(<-chan int, 2)

	//? the 1 , 2, tells how many channels to use or Create (Buffer)
	c := make(chan int, 1)

	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("------")
	fmt.Printf("%T\n", c)
}
