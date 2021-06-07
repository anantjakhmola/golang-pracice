package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		fmt.Println("hello world from one")
		wg.Done()
	}()
	go func() {
		fmt.Println("hello world from two")
		wg.Done()
	}()
	go func() {
		fmt.Println("hello world from three")
		wg.Done()
	}()
	go func() {
		fmt.Println("hello world from four")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())

}
