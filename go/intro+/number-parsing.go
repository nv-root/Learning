package main

import "fmt"
import "strconv"

func main() {

	numStr := "12345"
	num, err := strconv.Atoi(numStr) // string -> int

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", num)

	num32, err := strconv.ParseInt(numStr, 10, 32) // return int64, even after specifying the bitsize 32
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", num32)

	floatStr := "3.14159"
	floatNum, _ := strconv.ParseFloat(floatStr, 32) // same float64

	fmt.Printf("%T\n", floatNum)
	fmt.Printf("%.2f\n", floatNum)

	binStr := "1010"
	binNum, _ := strconv.ParseInt(binStr, 2, 8) // int64 bin -> int
	fmt.Printf("%T\n", binNum)
	fmt.Printf("%b\n", binNum)
	fmt.Println(binNum)

	hexStr := "ABC"
	hexNum, _ := strconv.ParseInt(hexStr, 16, 8) // int64 hex -> int
	fmt.Printf("%T\n", hexNum)
	fmt.Printf("%x\n", hexNum)

}
