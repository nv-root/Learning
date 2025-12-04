package main

import "fmt"
// import "os"

func process(pid int) {

	defer fmt.Println("diferred process", pid)
	defer fmt.Println("diferred process", pid)
	defer fmt.Println("diferred process", pid)
	if pid < 0 {
		fmt.Println("Before panic")
		panic("pid must be a non-negative number")
		// defer fmt.Println("diferred process")
	}
	fmt.Println("PID:", pid)
}

func process2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start process2")
	panic("Something went wrong")

}

func main() {
	// panic(interface{})

	defer fmt.Println("Defer statement in main()")
	fmt.Println("main() function started")

	// os.Exit(1) // bypasses defer statements, terminates immediately

	process(2)
	// process(-2)

	process2()
	process2()
	process2()

	fmt.Println("After function calls")

	fmt.Println("main() function exited")

}
