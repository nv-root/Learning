package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("This will run only once")
}

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine", i)
			once.Do(initialize)
		}()
	}

	wg.Wait()
}
