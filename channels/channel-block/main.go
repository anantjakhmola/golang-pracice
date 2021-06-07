//! REMEMBER CHANNELS BLOCK THERE MUST BE SOMETHING TO RECEIVE SOMETHING FROM THE CHANNEL

package main

import "fmt"

func main() {
	c := make(chan int)
	//c := make(chan int,1)

	go func() {
		c <- 42
		//c <- 33

	}()

	fmt.Println(<-c)
}
