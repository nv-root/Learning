package main

import (
	"fmt"
	"time"
)

func periodicTask() {
	fmt.Println("Performing periodic task at:", time.Now())
}

func tickerEx() {
	ticker := time.NewTicker(1 * time.Second)
	stop := time.After(5 * time.Second)

	defer func() {
		ticker.Stop()
		fmt.Println("Ticker stopped")
	}()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println(tick)
		case <-stop:
			fmt.Println("Stopping ticker...")
			return
		}
	}
}

func main() {
	tickerEx()
}
