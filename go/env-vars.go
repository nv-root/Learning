package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// user := os.Getenv("USER")
	// home := os.Getenv("HOME")
	//
	// fmt.Println(user)
	// fmt.Println(home)

	// err := os.Setenv("ENV", "development")
	// if err != nil {
	// 	fmt.Println("Error setting env var:", err)
	// }

	fmt.Println("ENV:", os.Getenv("ENV"))

	for _, e := range os.Environ() {
		fmt.Print(strings.SplitN(e, "=", 2)[0], "=")
		fmt.Println(strings.SplitN(e, "=", 2)[1])
	}

	err := os.Unsetenv("ENV")
	if err != nil {
		fmt.Println("Error unsetting env var - ENV", err)
		return
	}

	fmt.Println("Unset env var - ENV")
}
