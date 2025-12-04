// runes are int32 values of unicode points
package main

import "fmt"
import "unicode/utf8"

func main() {
	str := "Hello world"
	fmt.Println(str)
	fmt.Println(str[0])

	str += " world"
	fmt.Println(str)

	a := "apple"
	b := "banana"
	A := "Apple"
	App := "App"

	fmt.Println(a < b)
	fmt.Println(a < A)

	fmt.Println(A < App)
	fmt.Println(a < App)

	for i, char := range A{
		fmt.Printf("Character at index %d is %c - %v\n", i, char, char)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(a))


	var ch rune = 'a'
	fmt.Println(ch)
	fmt.Printf("%c\n", ch)
	fmt.Println(string(ch))

	fmt.Printf("%T, %T", ch, string(ch))

}
