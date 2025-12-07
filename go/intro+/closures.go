package main

import "fmt"

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

func main() {
	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())

	subtractor := func() func(int) int {
		fmt.Println("Initializing countdown")
		countdown := 100
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtractor(10))
	fmt.Println(subtractor(10))
	fmt.Println(subtractor(10))

}
