// if it quacks like a duck, it is a duck (or maybe not)
package main

import (
	"fmt"
	"math"
)

// interface
type Geometry interface {
	area() float64
	perimeter() float64
}

// Rectangle and Circle implement the interface Geometry
type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// ======================================
func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// =====================================
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) diameter() float64 {
	return 2 * c.radius
}

// =================================
type Rectangle2 struct {
	width, height float64
}

func (r Rectangle2) area() float64 {
	return r.width * r.height
}
// ===================================

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func myPrinter(i ...interface{}){ // can accept anything
	for _, v := range i{
		fmt.Printf("\n%v - %T", v, v)
	}
}


func main() {
	
	rect := Rectangle{width: 3, height: 4}
	circle := Circle{radius: 2}

	// can be accepted as Geometry - type inside measure()
	measure(rect)
	measure(circle)

	fmt.Println(circle.diameter())

	rect2 := Rectangle2{width: 1, height: 2}
	fmt.Println(rect2)
	// measure(rect2) // doesn't work - perimeter implementation missing


	myPrinter(1, 2, 3, "four", true)

}
