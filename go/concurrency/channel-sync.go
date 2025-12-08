package main

import (
	"fmt"
	"strconv"
	"time"
)

func chEx0() {
	ch := make(chan int)
	go func() {
		fmt.Println("Sending...")
		time.Sleep(time.Second)
		ch <- 9
		fmt.Println("Sent value")
	}()

	fmt.Println(<-ch)
}

func chEx1() {
	numGoroutines := 3
	done := make(chan int)

	for i := range numGoroutines {
		go func(id int) {
			fmt.Println("Goroutine", id+1, "working...")
			time.Sleep(time.Second)
			done <- id + 1
		}(i)
	}

	for range numGoroutines {
		<-done
	}
	fmt.Println("All Goroutines are finished")
}

// synchronizing data exchange
func chEx2() {
	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "Hello - " + strconv.Itoa(i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data) //signal completion by closing the channel
	}()

	for v := range data {
		time.Sleep(time.Second)
		fmt.Println("Received value:", v, time.Now())
	}

}

func main() {
	done := make(chan struct{}) // no data, just signaling
	go func() {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
		done <- struct{}{} //signaling completion
	}()

	<-done
	fmt.Println("Finished")

	chEx0()
	chEx1()
	// chEx2()
}
