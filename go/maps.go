package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int) // map[key-type]value-type
	fmt.Println(m)            // map[]

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

	//to check whether there is a value associated with the key
	value, present := m["k1"]
	fmt.Println(value, present)
	fmt.Println("map:", m)

	// map literal
	m2 := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("m2:", m2)

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// if maps.Equal(m, m2){
	// 	fmt.Println("m and m2 are equal")
	// }

	nestedMaps := make(map[int]map[string]string)

	nestedMaps[1] = map[string]string{"one": "ONE", "two": "TWO"}
	nestedMaps[2] = map[string]string{"one": "ONE", "two": "TWO", "three": "THREE"}

	fmt.Println(nestedMaps)
}
