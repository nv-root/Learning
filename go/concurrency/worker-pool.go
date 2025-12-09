package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second * 2)
		results <- job * 2
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}

}

func workerPool() {
	numWorkers := 3
	numJobs := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// creating workers
	for i := range numWorkers {
		go worker(i, jobs, results)
	}

	// sending jobs
	for i := range numJobs {
		jobs <- i

	}
	close(jobs)

	for range numJobs {
		fmt.Println("Result:", <-results)
	}
}

// =============================================
type ticketRequest struct {
	personID   int
	numTickets int
	cost       float64
}

func ticketProcesser(reqs <-chan ticketRequest, results chan<- int) {
	for req := range reqs {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %f \n", req.numTickets, req.personID, req.cost)
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func processTickets() {
	numRequests := 5
	price := 5.99
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int, numRequests)

	for range 3 {
		go ticketProcesser(ticketRequests, ticketResults)
	}

	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: float64(i+1) * price}
	}

	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully\n", <-ticketResults)
	}

}

func main() {
	// workerPool()

	processTickets()
}
