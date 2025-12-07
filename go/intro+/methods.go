package main

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

// we can associate methods for any type
func (i MyInt) IsPositive() bool {
	return i > 0
}

func (MyInt) Message() {
	fmt.Println("Welcome, Welcome")
}

func main() {

	rect := Rectangle{
		length: 10.42,
		width:  93.2,
	}

	fmt.Println(rect.length)
	fmt.Println(rect.width)
	fmt.Println("Area:", rect.Area())

	rect.Scale(2)

	fmt.Println("\nAfter scaling by 2")
	fmt.Println(rect.length)
	fmt.Println(rect.width)
	fmt.Println("Area:", rect.Area())

	myNum := MyInt(10)
	myNum.Message()
	fmt.Println(myNum.IsPositive())

	myNum = MyInt(-4)
	fmt.Println(myNum.IsPositive())

	s := Shape{
		Rectangle: Rectangle{
			length: 1,
			width:  2,
		},
	}

	fmt.Println(s.Area())

}
