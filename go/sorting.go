package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

type By func(p1, p2 *person) bool
type personSorter struct {
	ppl []person
	by  func(p1, p2 *person) bool
}

func (ps *personSorter) Len() int {
	return len(ps.ppl)
}

func (ps *personSorter) Less(i, j int) bool {
	return ps.by(&ps.ppl[i], &ps.ppl[j])
}

func (ps *personSorter) Swap(i, j int) {
	ps.ppl[i], ps.ppl[j] = ps.ppl[j], ps.ppl[i]
}

func (by By) Sort(ppl []person) {
	ps := &personSorter{
		ppl: ppl,
		by:  by,
	}
	sort.Sort(ps)
}

func main() {
	numbers := []int{5, 4, 3, 2, 1, 0}
	sort.Ints(numbers) // inplace
	fmt.Println("Sorted nums:", numbers)

	names := []string{"Zed", "William", "Abel", "Justin", "Ariana"}
	sort.Strings(names)
	fmt.Println("Sorted names:", names)

	ppl := []person{
		{"Ariana", 32},
		{"Justin", 30},
		{"Abel", 31},
	}
	fmt.Println("people:", ppl)

	age := func(p1, p2 *person) bool {
		return p1.age < p2.age
	}

	name := func(p1, p2 *person) bool {
		return p1.name < p2.name
	}

	lenName := func(p1, p2 *person) bool {
		return len(p1.name) < len(p2.name)
	}

	By(age).Sort(ppl)
	fmt.Println("Sorted by age:", ppl)

	By(name).Sort(ppl)
	fmt.Println("Sorted by name:", ppl)

	By(lenName).Sort(ppl)
	fmt.Println("Sorted by length of name:", ppl)

}
