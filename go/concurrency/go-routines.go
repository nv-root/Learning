// managed by the go runtime schuduler
// uses M:N  scheduling model
// M goroutines run on N os threads

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Entered sayHello()")
	time.Sleep(time.Second * 1)
	fmt.Println("Hello!")
}

func printNums() {
	for i := range 5 {
		fmt.Println("number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "Goroutines!" {
		fmt.Println("letter:", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)

	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in doWork()")
}

func main() {
	var err error

	fmt.Println("Entered main()")
	// executes in the background and returns to the main thread
	go sayHello() //Goroutines are non-blocking

	fmt.Println("After calling sayHello()")

	// execution order is not guaranteed
	go printLetters()
	go printNums()

	go func() { err = doWork() }()

	time.Sleep(time.Second * 2) // to keep the main thread alive

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Work completed")
	}
}
