package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker:", w.ID, "started", w.Task)
	time.Sleep(2 * time.Second)
	fmt.Println("Woker:", w.ID, "finished", w.Task)
}

func main() {

	var wg sync.WaitGroup
	tasks := []string{"task1", "task2", "task3"}

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	wg.Wait()
	fmt.Println("Construction finished")
}
