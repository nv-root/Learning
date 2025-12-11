package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	fmt.Println("PID:", os.Getpid())
	sigs := make(chan os.Signal, 1) // os.Signal is a type
	done := make(chan bool)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		done <- true
	}()

	// notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {

		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal")
				return
			default:
				fmt.Println("working...")
				time.Sleep(time.Second)

			}
		}

		// for sig := range sigs {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("\nInterrupted fu")
		// 	case syscall.SIGTERM:
		// 		fmt.Println("\nTerminated")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("\nHangup")
		// 	case syscall.SIGUSR1:
		// 		fmt.Println("user defined signal")
		// 		continue
		// 	}
		// 	fmt.Println("Graceful exit.")
		// 	os.Exit(0)
		// }
	}()

	for {
		time.Sleep(2 * time.Second)
	}
}
