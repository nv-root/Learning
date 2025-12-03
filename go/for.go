package main

import "fmt"
import "math/rand"
import "time"

func main() {

	// for working as while
	i := 1
	for i <= 3 {
		fmt.Println("for working as while", i)
		i++
	}

	// number guessing game
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("\n\nWelcome to the guessing game")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		if guess > target {
			fmt.Println("You guessed a little bigger")
		} else if guess < target {
			fmt.Println("You guessed a little smaller")
		} else {
			fmt.Println("Congratulations! You guessed it right.\n\n")
			break
		}
	}

	// for working as for
	for j := 0; j < 3; j++ {
		fmt.Println("for working as for", j)
	}

	// loop until break
	for {
		fmt.Println("loop until break")
		if i == 5 {
			break
		}

		i++
	}

	// range based
	for r := range 3 {
		fmt.Println("range", r)
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// iterating over a sequence
	numbers := []int{1, 2, 3, 4}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %v\n", index, value)
		fmt.Println(index, value)
	}

	// nested loops
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

}
