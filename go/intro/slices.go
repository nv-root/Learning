package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("Uninitialized:", s, s == nil, len(s) == 0)

	s = make([]string, 3, 4)
	// len(s) - no.of elements in the slice
	// cap(s) - how much a slice can grow before reallocation of the underlying array

	fmt.Println("Uninitialized:", s, s == nil, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// s[3] = "d" // panics - index out of range
	fmt.Println(s, len(s), cap(s))

	s = append(s, "d")
	fmt.Println(s, len(s), cap(s))

	// cap doubles if cap < 1024

	s = append(s, "d", "e", "f", "g", "h", "i", "j") // doubling is not sufficient . so ....
	fmt.Println(s, len(s), cap(s))                   // cap(s) = 11 ? go allocates the required length if doubling is not enough

	arr := [5]int{1, 2, 3, 4, 5}
	slce := arr[1:4]
	fmt.Println(slce)

	slceCopy := make([]int, len(slce))
	copy(slceCopy, slce)

	fmt.Println("slceCopy:", slceCopy)

	slceCopy[0] = 10001

	fmt.Println("slceCopy:", slceCopy)
	if slices.Equal(slce, slceCopy) {
		fmt.Println("slce and slceCopy are equal")
	}


	twoD := make([][]int, 3)
	for i := range 3{
		innerLength := i+1
		twoD[i] = make([]int, innerLength)
		for j := range innerLength{
			twoD[i][j] = i+j
		}
	}

	fmt.Println(twoD)

}
