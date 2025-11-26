package main

import "fmt"

func main(){

	// for working as while
	i := 1
	for i <= 3{
		fmt.Println("for working as while", i)
		i++
	}

	// for working as for
	for j := 0; j < 3; j++{
		fmt.Println("for working as for", j)
	}

	// loop until break
	for{
		fmt.Println("loop until break")
		if i == 5{
			break
		}

		i++
	}

	// range based
	for r := range 3{
		fmt.Println("range", r)
	}

	for n := range 6{
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}


}
