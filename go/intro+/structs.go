package main

import "fmt"

type Person struct {
	firstName   string
	lastName    string
	age         int
	address     Address
	PhoneNumber // embedded
}

type PhoneNumber struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

// value receiver
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// pointer receiver
func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       10,
		address: Address{
			city:    "LA",
			country: "USA",
		},
		PhoneNumber: PhoneNumber{
			home: "1234567890",
			cell: "0987654321",
		},
	}

	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
	fmt.Println(p.age)
	fmt.Println(p.fullName())
	p.incrementAgeByOne()
	p.incrementAgeByOne()
	fmt.Println(p.age)
	fmt.Println(p.address.country)
	fmt.Println(p.PhoneNumber.home)
	fmt.Println(p.home) // performs lookup in the embedded fields

	p2 := Person{
		firstName: "Justin",
		lastName:  "Bieber",
		age:       30,
	}

	fmt.Println(p2.fullName())

	// anonymous structs
	user := struct {
		username string
		email    string
	}{
		username: "user1",
		email:    "user1@gmail.com",
	}

	fmt.Println(user.username)
	fmt.Println(user.email)

	fmt.Println(p)

}
