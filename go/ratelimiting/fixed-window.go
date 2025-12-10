package main

import (
	"fmt"
	"sync"
	"time"
)

type FWRateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewFWRatelimiter(limit int, window time.Duration) *FWRateLimiter {
	return &FWRateLimiter{
		limit:     limit,
		window:    window,
		resetTime: time.Now().Add(window),
	}
}

func (rl *FWRateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	fmt.Println("count:", rl.count)

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
		fmt.Println("Window reset:", time.Now())
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {
	rateLimiter := NewFWRatelimiter(4, time.Second)
	fmt.Println("Initial Reset time:", rateLimiter.resetTime)
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			if rateLimiter.Allow() {
				fmt.Println("Request allowed")
			} else {
				fmt.Println("Request denied")
			}
			wg.Done()
		}()
		// time.Sleep(200 * time.Millisecond)
	}

	wg.Wait()
}
