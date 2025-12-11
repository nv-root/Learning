// helps in working with generics programs and types at runtime
// use cases - data serialization, generic types
package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	age  int
}

type greeter struct{}

func (g greeter) Greet(fname string, lname string) string {
	return "Hello " + fname + " " + lname
}

func reflectEx3() {
	g := greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)
	var methods []reflect.Method

	fmt.Println("Type of g:", t)

	for i := range t.NumMethod() {
		methods = append(methods, t.Method(i))
	}

	m := v.MethodByName(methods[0].Name)
	results := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("Doe")})

	for _, v := range results {
		fmt.Println(v)
	}

}

func reflectEx2() { // working with structs
	p1 := person{Name: "Alice", age: 30}
	v := reflect.ValueOf(p1)

	fmt.Println(p1)
	for i := range v.NumField() {
		fmt.Println("Field -", i, v.Field(i))
	}

	v1 := reflect.ValueOf(&p1).Elem()
	nameField := v1.FieldByName("Name")

	if nameField.CanSet() { // fields should be visible i.e uppercase starting
		nameField.SetString("Justin")
	} else {
		fmt.Println("Cannot set")
	}
	fmt.Println(p1)
}

func reflectEx1() {
	x := "str"
	v := reflect.ValueOf(x)

	t := v.Type()
	fmt.Println("Value:", v)
	fmt.Println("Type:", t)
	fmt.Println("Kind is String:", t.Kind() == reflect.String)
	fmt.Println("Is zero?", v.IsZero())

	y := 10
	v1 := reflect.ValueOf(&y).Elem()
	v2 := reflect.ValueOf(&y)

	fmt.Println("v2 type:", v2.Type())

	fmt.Println("Original value:", v1.Int())
	v1.SetInt(3)
	fmt.Println("Modified value:", v1.Int())

	//interface
	var intrfce interface{} = "Hello"
	v3 := reflect.ValueOf(intrfce)

	fmt.Println("v3 type:", v3.Type())
	if v3.Kind() == reflect.String {
		fmt.Println("string value", v3.String())
	}
}

func main() {

	reflectEx1()
	reflectEx2()
	reflectEx3()

}
