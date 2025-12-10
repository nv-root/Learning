package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(capacity int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, capacity),
		refillTime: refillTime,
	}
	for range capacity {
		rl.tokens <- struct{}{}
	}

	go rl.StartRefill()

	return rl
}

func (rl *RateLimiter) StartRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
				fmt.Println("Added one token at:", time.Now().Format("2006-01-02 15:04:05.0000"))
			default:
				fmt.Println("Bucket full")
			}
		}
	}

}

func (rl *RateLimiter) Allow() bool {
	fmt.Printf("New request at: %v", time.Now().Format("2006-01-02 15:04:05.0000"))
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, time.Second)

	fmt.Println("Initial tokens:", len(rateLimiter.tokens))
	for range 30 {
		if rateLimiter.Allow() {
			fmt.Println(" - Request allowed - Left: ", len(rateLimiter.tokens))
		} else {
			fmt.Println(" - Request denied - Left:", len(rateLimiter.tokens))
		}
		time.Sleep(200 * time.Millisecond)
	}
}
