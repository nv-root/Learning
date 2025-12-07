package main

import "fmt"

func main() {
	i := 10_000
	s := "hello"
	f := 3.14

	fmt.Printf("i - %v - %#v - %T - %v%%\n", i, i, i, i)
	fmt.Printf("s - %v - %#v - %T\n", s, s, s)
	fmt.Printf("f - %v - %#v - %T\n", f, f, f)

	i = 12
	// binary - decimal - with sign - octa -  - hex
	fmt.Printf("i - %b - %d - %+d - %o - %O - %x\n", i, i, i, i, i, i)

	// string - string with quotes - padding with spaces - - hex -
	fmt.Printf("s - %s - %q - %8s - %-8s - %x - % x\n", s, s, s, s, s, s)

	tb := true
	fb := false
	fmt.Printf("t - %t - %v\n", tb, tb)
	fmt.Printf("f - %t - %v\n", fb, fb)

}
