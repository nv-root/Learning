package main

import (
	"fmt"
	"time"
)

func timerEx() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for timer")

	// stopping the timer
	stopped := timer.Stop()
	if stopped {
		fmt.Println("Timer stopped")
	}

	timer.Reset(time.Second)
	fmt.Println("Timer reset")
	<-timer.C // blocking
	fmt.Println("Timer expired")
}

func longRunningOp() {
	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func timeoutEx() {
	timeout := time.After(2 * time.Second)
	done := make(chan bool)
	go func() {
		longRunningOp()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("Operation timedout")
	case <-done:
		fmt.Println("Operation completed")
	}
}

func delayedOp() {
	timer := time.NewTimer(2 * time.Second)
	done := make(chan bool)
	go func() {
		<-timer.C
		fmt.Println("Delayed operation executed")
		done <- true
	}()

	fmt.Println("Waiting for delayed operation to complete...")
	<-done
	fmt.Println("Delayed operation completed")
}

func multipleTimers() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("Timer1 completed")
	case <-timer2.C:
		fmt.Println("Timer2 completed")
	}
}

func main() {
	fmt.Println("Entered main")
	// timerEx()

	// timeoutEx()

	// delayedOp()

	multipleTimers()

}
