package main

import (
	"fmt"
	"time"
)

type statefulWorker struct {
	count int
	ch    chan int
}

func (w *statefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current count:", w.count)
			}
		}
	}()
}

func (w *statefulWorker) Send(value int) {
	w.ch <- value
}

func main() {

	stWorker := statefulWorker{
		ch: make(chan int),
	}

	stWorker.Start()

	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}

}
