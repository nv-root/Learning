package main

import "fmt"
import "unsafe"

const S string = "constant"

func main() {

	fmt.Println(S)
	// S = "reassigned"
	fmt.Printf("Type of S: %T\n", S)

	const C = 10
	fmt.Println(C)
	fmt.Printf("Type of C: %T\n", C)

	const X = 500
	fmt.Printf("Type of X: %T\n", X)

	const (
		MONDAY    = 1
		TUESDAY   = 2
		WEDNESDAY = 3
	)

	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)

	fmt.Println(unsafe.Sizeof(int(0)))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int32(0)))
	fmt.Println(unsafe.Sizeof(int64(0)))

	var i int = 64
	var i64 int64 = 64
	// fmt.Println(i == i64)
	fmt.Println(int64(i) == i64)

}
