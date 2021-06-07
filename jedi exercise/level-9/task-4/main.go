package main

import (
	"fmt"
	"sync"
)

func main() {

	var waitgroup sync.WaitGroup

	incrementor := 0
	goroutines := 100

	waitgroup.Add(goroutines)
	var m sync.Mutex

	for i := 0; i < goroutines; i++ {
		go func() {
			m.Lock()
			v := incrementor
			//runtime.Gosched()
			v++
			incrementor = v
			fmt.Println(incrementor)
			m.Unlock()

			waitgroup.Done()
		}()

	}
	waitgroup.Wait()
	fmt.Println("end value", incrementor)

}
