package main

import "fmt"

func process(i int) {
	// defer - schedules a function call to be executed just before the surrounding function exits
	// only "executed" after,  "evaluated" immediately 
	// LIFO
	defer fmt.Println("diferred i",i)
	defer fmt.Println("first diferred statement executed")
	defer fmt.Println("second diferred statement executed")
	defer fmt.Println("third diferred statement executed")
	fmt.Println("Normal statement executed")

	i++
	fmt.Println(i)
}

func main() {

	process(100)

}
