package main

import (
	"fmt"
	"sync"
)

var (
	rwmu      sync.RWMutex
	rwCounter int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	fmt.Println("Read counter:", rwCounter)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwmu.Lock()
	rwCounter = value
	fmt.Println("Written value", value, "for counter")
	rwmu.Unlock()
}

func main() {

	var wg sync.WaitGroup

	// wg.Add(1)
	// go writeCounter(&wg, 18)
	// wg.Wait()

	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	go writeCounter(&wg, 18)
	wg.Wait()

}
