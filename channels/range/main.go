package main

import "fmt"

func main() {

	c := make(chan int)
	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//receiver
	//? Range will pull of from a channel until it is closed

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("bout to exit")

}
