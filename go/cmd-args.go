package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// for i, arg := range os.Args {
	// 	fmt.Println("Arg", i, ":", arg)
	// }
	//
	// var name string
	// var age int
	//
	// flag.StringVar(&name, "name", "nv", "Name of the user")
	// flag.IntVar(&age, "age", 18, "Age of the user")
	//
	// flag.Parse()
	//
	// fmt.Println(name)
	// fmt.Println(age)
	//
	//sub commands
	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subcommand1.Bool("processing", false, "Command processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subcommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("SubCommand1:")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("SubCommand2:")
		fmt.Println("language:", *flagsc2)
	default:
		fmt.Println("No subcommand entered")
		os.Exit(1)

	}

}
