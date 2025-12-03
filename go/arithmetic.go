package main

import "fmt"
import "math"

func main() {

	var a, b int = 5, 2.0
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	// integer / integer = integer
	// integer / float = float
	// float / integer = float
	var result2 float64 = 22.0 / 7
	fmt.Println(result2)


	// overflow with ints - signed and unsigned
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt += 1
	fmt.Println(maxInt)
	
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt += 1
	fmt.Println(uMaxInt)


	// underflow with floats
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat/math.MaxFloat64
	fmt.Println(smallFloat)


}
