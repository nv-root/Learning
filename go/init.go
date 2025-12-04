package main

import "fmt"

// executed by go for initializing the packages
func init() {
	fmt.Println("Initializing package... 1")
}


func init() {
	fmt.Println("Initializing package... 2")
}


func init() {
	fmt.Println("Initializing package... 3")
}


func main() {
	fmt.Println("Inside the main()")
}
