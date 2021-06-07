package main

import "fmt"

func main() {
	//? using bufferend channel
	//c := make(chan int, 1)

	c := make(chan int)
	//using an anonynmous function
	go func() {
		c <- 42

	}()

	fmt.Println(<-c)
}
