package main

import "fmt"

func main() {

	num := 148
	fmt.Printf("%05d\n", num)

	message := "Hello"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	message2 := `Hello \nWorld`
	message3 := "Hello \nWorld"
	fmt.Println(message2)
	fmt.Println(message3)

}
