// demonstrates field promotion
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	empId  string
	salary float64
}

func (p Person) introduce(){
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.name, p.age)
}

// overrides the above inroduce()
// outer type (Employee) method overriding the inner type(Person) method
func (e Employee) introduce(){
	fmt.Printf("Employee ID: %s, Salary: %.2f\n", e.empId, e.salary)
}

func main() {
	emp := Employee{
		Person: Person{name: "john", age: 32},
		empId:  "E001",
		salary: 50000,
	}

	fmt.Println(emp)
	fmt.Println(emp.name)
	fmt.Println(emp.age)
	fmt.Println(emp.empId)
	emp.introduce()
}
