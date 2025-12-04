package main

import "fmt"

func zeroval(ival int){
	ival = 0
}

func zeroptr(iptr *int){
	fmt.Println("before modifying", *iptr)
	*iptr = 0
	fmt.Println(iptr)
	fmt.Println("after modifying", *iptr)
}

func main(){

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("i address:", &i)
}
