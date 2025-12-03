package main

import "fmt"
import "time"

func main() {
	var num int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&num)

	// if-else
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// else-if
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// nested if
	num2 := 6
	if num2%2 == 0 {
		if num2%3 == 0 {
			fmt.Println("Divisible by both 2 and 3")
		} else {
			fmt.Println("Only divisible by only 2")
		}
	} else {
		fmt.Println("Not divisible by 2")
	}

	//switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// switch with multiple expressions in case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("DAY:", time.Now().Weekday(), time.Now())
		fmt.Println("It's the weekend")
	default:
		fmt.Println("DAY:", time.Now().Weekday(), time.Now())
		fmt.Println("It's a weekday")
	}

	// switch without expression
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	num3 := 2
	switch {
	case num3 > 1:
		fmt.Println("Greater than 1")
		fallthrough // checks another case too
	case num3 == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not 2")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(323)
	whatAmI("hey")

}
