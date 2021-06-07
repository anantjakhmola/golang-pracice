package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var waitgroup sync.WaitGroup

	incrementor := 0
	goroutines := 100

	waitgroup.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			v := incrementor
			runtime.Gosched()
			v++
			incrementor = v
			fmt.Println(incrementor)

			waitgroup.Done()
		}()

	}
	waitgroup.Wait()
	fmt.Println("end value", incrementor)

}
