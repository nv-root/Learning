// sync mechanism for goroutines -> "communication is synchronization" :)
package main

import (
	"fmt"
	"time"
)

func main() {
	// creating an unbuffered channel (capacity = 0)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		// unbuffered channel -> needs an immediate receiver
		// blocks until there is a receiver ready to take the value
		greeting <- greetString
		greeting <- "World"

		for _, char := range "channels" {
			greeting <- "Alphabet: " + string(char)
		}
	}()

	// receiver is part of the main Goroutine if not wrapped in a separate goroutine
	go func() {
		// blocks until a sender sends
		receiver := <-greeting
		fmt.Println(receiver)

		receiver = <-greeting
		fmt.Println(receiver)

		for range 8 {
			receiver = <-greeting
			fmt.Println(receiver)
		}

	}()

	// should not use this shit if prod
	time.Sleep(1 * time.Second)
	fmt.Println("EOF")
}
