package main

import "fmt"

func main(){
	nums := []int{10, 20, 30}
	sum := 0

	fmt.Println("slice:", nums)

	for _, num := range nums{
		sum += num
	}

	fmt.Println("sum: ", sum)

	for i, num := range nums{
			fmt.Println("index:", i, "num:", num)
	}


	kvs := map[int]string{1:"one", 2:"two", 3:"three", 4:"four"}
	fmt.Println(kvs)
	for k, v := range kvs{
		fmt.Println("key:", k, "value:", v)
	}

	// i - starting byte index of the rune, c - rune
	for i, c := range "go-go-go"{ 
		// fmt.Println(i, c)
		fmt.Printf("\nIndex: %d, Rune: %c\n", i, c)
	}
}
