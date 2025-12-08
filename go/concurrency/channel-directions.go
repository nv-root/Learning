package main

import "fmt"

// send-only channel
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// receive-only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("consumer():", value)
	}
}

func main() {

	ch := make(chan int)

	producer(ch)
	consumer(ch)

	// for value := range ch {
	// 	fmt.Println("Received:", value)
	// }

}
