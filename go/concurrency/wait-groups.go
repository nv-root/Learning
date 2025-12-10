package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrements the semaphore
	fmt.Println("Worker", id, "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker", id, "finished")
}

func workerT2(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker ID", id, "starting...")
	time.Sleep(time.Second)
	for task := range tasks {
		results <- task
	}
	fmt.Println("Worker ID", id, "finished")
}

func waitGroupEx1() {
	var wg sync.WaitGroup // counting semaphore
	numWorkers := 3

	// wg.Add(numWorkers) // no.of workers

	for i := range numWorkers {
		go worker(i, &wg)
	}

	wg.Wait() // blocks until wg becomes zero
	fmt.Println("All workers finished")
}

func waitGroupEx2() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 6

	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go workerT2(i, tasks, results, &wg)
	}

	// sending tasks
	for i := range numJobs {
		tasks <- i
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}

	fmt.Println("All T2Workers finished")

}

// ===========================================================

func main() {
	waitGroupEx1()

	// waitGroupEx2()
}
