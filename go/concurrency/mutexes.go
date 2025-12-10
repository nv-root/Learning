// mutexes (mutual exclusion) - a sync mechanism to prevent multiple routines from accessing shared resouces simultaneously

// why? - data integrity, synchronization, avoiding race conditions
// Lock() - acquire lock and protects, if not blocking
// Unlock() - releases lock
// TryLock() - attempts to acquire lock without blocking

// - to check data races -> `go run -race mutexes.go`
package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu    sync.Mutex
	value int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}

func mutexExample1() {
	var wg sync.WaitGroup
	counter := &counter{}
	numGoroutines := 10

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// counter.value++ // results in race condition
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter:", counter.getValue())
}

func muxtexExample2() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 10
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++ // critical section -> between lock() and unlock()
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Println("Final counter 2:", counter)

}

func main() {
	// mutexExample1()

	muxtexExample2()

}
