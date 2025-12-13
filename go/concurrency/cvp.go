package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func printNums() {
	for i := range 5 {
		fmt.Println(i, time.Now())
	}
}

func printLetters() {
	for _, letter := range "ABCDE" {
		fmt.Println(string(letter), time.Now())
	}
}

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id, "is starting -", time.Now())
	for range 100_000_000 {
	}
	fmt.Println("Task", id, "finished -", time.Now())
}

func main() {
	// go printNums()
	// go printLetters()
	//
	// time.Sleep(2 * time.Second)

	numThreads := 4
	runtime.GOMAXPROCS(numThreads)
	var wg sync.WaitGroup

	for i := range numThreads {
		wg.Add(1)
		go heavyTask(i, &wg)
	}

	wg.Wait()

}
