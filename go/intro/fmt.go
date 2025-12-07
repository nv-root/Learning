package main

import "fmt"

func main(){

	// printing functions
	fmt.Print("Hello ")
	fmt.Print("World ")
	fmt.Print(2.141)

	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println(2.141)

	fmt.Printf("%s %d\n", "john", 18)
	fmt.Printf("%x %b\n", "john", 18)


	// formatting functions - return
	s := fmt.Sprint("hello ", "world") // concatenates
	fmt.Print(s)
	
	s = fmt.Sprintln("hello", "world") // adds spaces and new line
	fmt.Print(s)
	fmt.Print(s)

	sf := fmt.Sprintf("%s %d", "john", 18)
	fmt.Print(sf)
	fmt.Print(sf)

	// scanning function
	var name string
	var age int
	fmt.Printf("\nEnter name and age: ")
	fmt.Scan(&name, &age)
	fmt.Println(name, age)	

	fmt.Printf("\nEnter name: ")
	fmt.Scanln(&name)
	fmt.Println(name)

	fmt.Printf("Enter name and age : ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println(name, age)


	// error formatting
	err := checkAge(15)
	if err != nil{
		fmt.Println(err)
	}
	
}

func checkAge(age int) error{
	if age < 18{
		return fmt.Errorf("Age %d is too young to drive", age)
	}

	return nil
}
