package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func reverseString(str string) string {
	if len(str) == 0 {
		return ""
	}

	return reverseString(str[1:]) + str[0:1]
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}

func main() {

	var n int
	fmt.Printf("Enter a number to find the factorial: ")
	fmt.Scan(&n)
	fmt.Println(fact(n))

	fmt.Println(reverseString("hello"))
	fmt.Println(reverseString("blah-blah-blah"))

	fmt.Println(sumOfDigits(439))
	fmt.Println(sumOfDigits(9))

}
