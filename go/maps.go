package main

import "fmt"

func main(){

	m := make(map[string]int) // map[key-type]value-type

	m["k1"] = 10
	m["k2"] = 20

	fmt.Println("map:", m)

	fmt.Println(m["k1"])
	fmt.Println(m["k2"])
	fmt.Println(m["k3"]) // 0?

	fmt.Println("len:", len(m))
	
	delete(m, "k2") // delete k2
	fmt.Println("map:", m)

	clear(m) // remove all pairs
	fmt.Println("map:", m)

	m["k1"] = 100
	value, present := m["k1"]
	fmt.Println(value, present)
	fmt.Println("map:", m)


	m2 := map[int]string{1:"one", 2:"two", 3:"three"}
	fmt.Println("m2:",m2)


}
