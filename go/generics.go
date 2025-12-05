package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {

	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println("Stack elements: ")
	for _, element := range s.elements {
		fmt.Println(element)
	}
}

func main() {
	x, y := 1, 2
	fmt.Println(x, y)

	x, y = swap(x, y)
	fmt.Println(x, y)

	intStack := Stack[int]{}

	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.push(4)
	intStack.push(5)
	intStack.printAll()

	fmt.Println(intStack.pop())
	intStack.printAll()

	strStack := Stack[string]{}

	strStack.push("apple")
	strStack.push("banana")
	strStack.push("carrot")
	strStack.printAll()

}
