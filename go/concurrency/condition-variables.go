// waits for a specific state
// until then..., sleep

// Wait() -> puts the caller to sleep and temporarily unlocks
// Signal() -> wakes up other goroutines
// Broadcast()

package main

import (
	"fmt"
	"sync"
	"time"
)

const bufSize = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// buffer is full -> wait, consumer can proceed to consume
	for len(b.items) == bufSize {
		b.cond.Wait()
	}
	b.items = append(b.items, item)
	fmt.Println("Produced:", item)

	// signaling after producing
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	// buffer is empty -> wait, producer con proceed to produce
	for len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item)

	// signaling after consuming so the producer can continue to add
	b.cond.Signal()
	return item
}

func cvProducer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range 10 {
		b.produce((i + 1) * 10)
		time.Sleep(time.Millisecond)
	}
}

func cvConsumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 10 {
		b.consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {

	buf := newBuffer(bufSize)

	var wg sync.WaitGroup

	wg.Add(2)
	go cvProducer(buf, &wg)
	go cvConsumer(buf, &wg)

	wg.Wait()

}
