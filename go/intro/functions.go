// public functions - uppercase starting letter - can be used in other packages
// private functions - lowercase

package main

import "fmt"
import "errors"

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

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	// go returns named return values without explicit return
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("a and b are equal")
	}
}

// variadic functions - functions that can be called with any no.of arguments
func sumNums(returnString string, nums ...int) {
	// nums is an array
	fmt.Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(returnString, total)
}

func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {

	res := add(10, 20)
	fmt.Println("10+20 =", res)

	fmt.Println("10+20+30 =", sum(10, 20, 30))

	id, name := vals()
	fmt.Println(id, name)

	_, MyName := vals()
	fmt.Println(MyName)

	sumNums("sum of 1, 2, 3 is", 1, 2, 3)
	nums := []int{1, 2, 3, 4, 5}
	sumNums("sum of 1, 2, 3, 4, 5 is", nums...)

	// anonymous functions
	greet := func(name string) {
		fmt.Println("Hello", name)
	}

	greet(name)

	// assigning a function
	greet2 := greet

	greet2(name)

	result := applyOperation(10, 20, add)
	fmt.Println(result)

	multiplyBy3 := createMultiplier(3)
	fmt.Println(multiplyBy3(2))

	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)

	compareResult, err := compare(10, 10)
	if err == nil {
		fmt.Println(compareResult)
	} else {
		fmt.Println("Error:", err)
	}

}
