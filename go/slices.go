package main

import "fmt"

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
	fmt.Println(s, len(s), cap(s)) // cap(s) = 11 ? go allocates the required length if doubling is not enough 
}
