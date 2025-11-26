package main

import "fmt"
import "unsafe"

const s string = "constant"

func main(){

	fmt.Println(s)
	// s = "reassigned"
	fmt.Printf("Type of s: %T\n", s)

	const c = 10
	fmt.Println(c)
	fmt.Printf("Type of c: %T\n", c)

	const x  = 500
	fmt.Printf("Type of x: %T\n", x)

	fmt.Println(unsafe.Sizeof(int(0)))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int32(0)))
	fmt.Println(unsafe.Sizeof(int64(0)))

	var i int = 64
	var i64 int64 = 64
	// fmt.Println(i == i64)
	fmt.Println(int64(i) == i64)



}
