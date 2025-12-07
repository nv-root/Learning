package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	City    string `json:"city"` // struct tags
	Country string `json:"country"`
}

type Person struct {
	Name         string `json:"name"`
	Age          int    `json:"age,omitempty"` //  if zero_value ? omit
	EmailAddress string `json:"email,omitempty"`
	// Address      Address `json:"address"`
	Address Address `json:"-"` // always omit with "-"
}

func main() {
	person := Person{Name: "NV"}
	person = Person{Name: "Justin", Age: 30, EmailAddress: "justin@gmail.com", Address: Address{City: "Toronto", Country: "Canada"}}

	//encoding
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	//decoding
	var decodedPerson Person
	err = json.Unmarshal([]byte(jsonData), &decodedPerson)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(person)
	fmt.Println(decodedPerson)

	// list of objects
	listOfAddresses := []Address{
		{City: "NY", Country: "US"},
		{City: "LA", Country: "US"},
		{City: "California", Country: "US"},
	}

	jsonList, err := json.Marshal(listOfAddresses)
	if err != nil {
		log.Fatal("Error marshalling to JSON:", err)
	}

	fmt.Println(listOfAddresses)
	fmt.Println(string(jsonList))

	//working with unknown types
	unknownJsonData := `{"one":"ONE", "two": 1, "three" :"THREE", "address":{"country":"India"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(unknownJsonData), &data)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	fmt.Println(data["address"].(map[string]interface{})["country"])

}
