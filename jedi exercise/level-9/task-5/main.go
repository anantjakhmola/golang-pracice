//We will use atomic to get rid of the race condititon
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var waitgroup sync.WaitGroup
	var incrementor int64

	//incrementor := 0
	goroutines := 100

	waitgroup.Add(goroutines)
	//var m sync.Mutex

	for i := 0; i < goroutines; i++ {
		go func() {
			//m.Lock()
			atomic.AddInt64(&incrementor, 1)
			//incrementor
			//runtime.Gosched()
			//v++
			//incrementor = v
			//m.Unlock()
			fmt.Println(atomic.LoadInt64(&incrementor))

			waitgroup.Done()
		}()

	}
	waitgroup.Wait()
	fmt.Println("end value", incrementor)

}
