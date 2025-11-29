package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sum(a, b, c int) int { // omitting type specification for consecutive params of the same type until the last one :)
	return a + b + c
}

// multiple return values
func vals() (int, string) {
	return 101, "nv-root"
}

// variadic functions - functions that can be called with any no.of arguments
func sumNums(nums ...int) {
	// nums is an array
	fmt.Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {

	res := add(10, 20)
	fmt.Println("10+20 =", res)

	fmt.Println("10+20+30 =", sum(10, 20, 30))

	id, name := vals()
	fmt.Println(id, name)

	_, MyName := vals()
	fmt.Println(MyName)

	sumNums(1, 2, 3)
	nums := []int{1, 2, 3, 4, 5}
	sumNums(nums...)

}
