// atomic - indivisible, uninterruptable operations without intermidiary state

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	count int64
}

func (ac *atomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *atomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func acEx() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &atomicCounter{}

	// value := 0

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// value++
			}

		}()
	}

	wg.Wait()
	fmt.Println("final counter value:", counter.getValue())
	// fmt.Println(value)
}

func main() {

	acEx()

}
