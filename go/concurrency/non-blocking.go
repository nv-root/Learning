package main

import (
	"fmt"
	"time"
)

//
// func producer(ch chan<- int) {
// 	for i := range 5 {
// 		ch <- i
// 	}
//
// 	close(ch)
// }
//
// func filter(in <-chan int, out chan<- int) {
// 	for val := range in {
// 		if val%2 == 0 {
// 			out <- val
// 		}
// 	}
// 	close(out)
// }
//
// func chClose() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
//
// 	go producer(ch1)
// 	go filter(ch1, ch2)
//
// 	for val := range ch2 {
// 		fmt.Println("Even:", val)
// 	}
// }

func main() {

	ch := make(chan int)

	// non-blocking receive operation
	select {
	case m := <-ch:
		fmt.Println("Received:", m)
	default:
		fmt.Println("No messages available")
	}

	// non-blocking send operation
	select {
	case ch <- 1:
		fmt.Println("Sent message")
	default:
		fmt.Println("Channel is not ready to receive")
	}

	data := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received:", d)
			case <-quit:
				fmt.Println("Exiting")
				return
			default:
				fmt.Println("Waiting for data")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}
	quit <- true

	// chClose()

}
