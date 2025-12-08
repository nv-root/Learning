// again, using time.Sleep() is a very shitty practice
package main

import (
	"fmt"
	"time"
)

// =========================================================
func multiplexing() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	time.Sleep(2 * time.Second)

	for range 2 {
		select {
		case r := <-ch1:
			fmt.Println("received from ch1:", r)
		case r := <-ch2:
			fmt.Println("received from ch2:", r)
		default:
			fmt.Println("No channel is ready")
		}
	}
}

// =========================================================
func timeout() {
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
		close(ch)
	}()

	select {
	case m := <-ch:
		fmt.Println("Received:", m)
	case <-time.After(2 * time.Second):
		fmt.Println("Timedout")
	}
}

// =========================================================
func chClose() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 100
		close(ch) // signals that no more values will be sent
	}()

	for {
		select {
		case m, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received:", m)
		}
	}
}

func main() {
	fmt.Println("\nMultiplexing demo")
	multiplexing()

	fmt.Println("\nTimeout demo")
	timeout()
	fmt.Println("EOF")

	chClose()
}
