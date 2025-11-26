package main

import "fmt"

var outside = "OUTSIDE"
// outside2 := "OUTSIDE2"

func main(){

	var name = "nv-root"
	fmt.Println("Name:",name)

	var age, language = 18, "go"
	fmt.Println("Age:", age, "\nLanguage:", language)

	var isStudent = true
	fmt.Println("Is Student?", isStudent)

	var emptyIntVar int
	var emptyStrVar string
	fmt.Println(emptyIntVar)
	fmt.Println("-_-", emptyStrVar)

	fruit := "apple"
	fmt.Println(fruit)

	outside2 := "OUTSIDE2" // shorthand notation can be only used inside a function

	fmt.Println(outside)
	fmt.Println(outside2)

}
