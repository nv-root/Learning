package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type person struct {
	Name  string `xml:"name,omitempty"`
	Age   int    `xml:"-"`
	Email string `xml:"email"`
}

type book struct {
	XMLName xml.Name `xml:"book"`
	Color   string   `xml:"color,attr"`
	Author  string   `xml:"author,attr"`
	Title   string   `xml:"title,attr"`
}

func main() {

	person1 := person{Name: "", Age: 30, Email: "justin@gmail.com"}

	xmlData, err := xml.Marshal(person1)
	if err != nil {
		log.Fatal("Error marshalling data into xml:", err)
	}

	fmt.Println(string(xmlData))

	xmlData1, err := xml.MarshalIndent(person1, "", "  ")
	if err != nil {
		log.Fatal("Error indenting xml data:", err)
	}

	fmt.Println(string(xmlData1))

	// unmarshalling
	xmlRaw := `<person><name>Justin</name><age>30</age><email>justin@gmail.com</email></person>`

	var person2 person

	err = xml.Unmarshal([]byte(xmlRaw), &person2)
	if err != nil {
		log.Fatalln("Error unmarshalling xml:", err)
	}

	fmt.Println(person1)
	fmt.Println(person2)

	book1 := book{Color: "blue", Author: "nv", Title: "idk"}

	xmlData2, _ := xml.MarshalIndent(book1, "", "  ")
	fmt.Println(string(xmlData2))

}
