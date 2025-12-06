package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	val := rand.New(rand.NewSource(time.Now().Unix())) // fixed random number with seeding
	// fmt.Println(rand.Intn(100) + 1)
	fmt.Println(val.Intn(287000))

	fmt.Println(math.Floor(rand.Float64() * 10))

	for {
		fmt.Println("Welcom to dice game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")

		var choice int

		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for playing.")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Printf("You rolled a %d and %d\n", die1, die2)
		fmt.Println("Total: ", die1+die2)

	}

}
