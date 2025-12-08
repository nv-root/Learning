package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	go func() {
		// fmt.Println("Entered goroutine")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
	}()

	fmt.Println("Blocking starts")
	// blocking on send if the buffer is full
	ch <- 3
	fmt.Println("Blocking ends")

	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	fmt.Println("Buffered channels")

}

// func main() {
//
// 	ch := make(chan int, 2)
//
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()
//
// 	fmt.Println("Value:", <-ch)
// 	fmt.Println("Value:", <-ch)
//
// 	fmt.Println("EOF")
// }
