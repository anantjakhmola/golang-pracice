package main

import "fmt"

func main() {

	c := make(chan int)

	//send
	go foo(c)
	//receiver
	bar(c)

	fmt.Println("bout to exit")

}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)

}
