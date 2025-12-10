package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	queue    chan int
	interval time.Duration
	done     chan struct{}
}

func NewLeakyBucket(capacity int, interval time.Duration) *LeakyBucket {
	return &LeakyBucket{
		queue:    make(chan int, capacity),
		interval: interval,
		done:     make(chan struct{}),
	}
}

func (lb *LeakyBucket) AddRequest(id int) bool {
	select {
	case lb.queue <- id:
		return true // added
	default:
		return false // dropped
	}
}

func (lb *LeakyBucket) Start(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(lb.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C: // leak at a fixed interval
				select {
				case req := <-lb.queue:
					fmt.Println("Processing:", req)
				default:
					// nothing to do
				}
			case <-lb.done:
				fmt.Println("Stopped.")
				return
			}
		}
	}()
}

func (lb *LeakyBucket) Stop() {
	close(lb.done)
}

func main() {

	leakyBucket := NewLeakyBucket(5, 1*time.Second)
	var wg sync.WaitGroup

	wg.Add(1)
	leakyBucket.Start(&wg)

	defer wg.Wait()
	defer leakyBucket.Stop()

	for i := range 10 {
		if leakyBucket.AddRequest(i + 1) {
			fmt.Println("Request accepted:", i+1)
		} else {
			fmt.Println("Dropped:", i+1)
			time.Sleep(200 * time.Millisecond)
		}
	}

	time.Sleep(3 * time.Second)
	for i := range 10 {
		if leakyBucket.AddRequest(i + 1) {
			fmt.Println("Request accepted:", i+1)
		} else {
			fmt.Println("Dropped:", i+1)
			time.Sleep(500 * time.Millisecond)
		}
	}

	time.Sleep(10 * time.Second)
	fmt.Println("All done")

}
