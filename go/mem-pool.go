// for reusing memory
// reduced GC overhead
package main

import (
	"fmt"
	"sync"
)

type poolPerson struct {
	name string
	age  int
}

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new person...")
			return &poolPerson{}
		},
	}

	person1 := pool.Get().(*poolPerson)
	person1.name = "John"
	person1.age = 20

	fmt.Println("person1:", *person1)
	// now pool has the reference to person1
	// must assume garbage data, even if it contains pervious data
	// do not use objects after returning to the pool, overwrite for reusing
	pool.Put(person1)

	// contains pervious data
	// must not be used
	fmt.Println("person1 after putting:", *person1)

	person2 := pool.Get().(*poolPerson)
	// now person2 has reference to person1
	fmt.Println("person2:", *person2)
	//overwrite for reusing memory
	person2.name = "Abel"
	person2.age = 32

	fmt.Println("person2:", *person2)

	fmt.Println("person1 after overwritten by person2:", *person1)

	person3 := pool.Get().(*poolPerson)
	fmt.Println("person3:", person3)

	person4 := pool.Get().(*poolPerson)
	fmt.Println("person4:", person4)
}
