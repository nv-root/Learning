package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	// tmplt, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	//
	// if err != nil{
	// 	panic(err)
	// }

	tmplt := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing\n"))

	// data := map[string]interface{}{
	// 	"name" :"John",
	// }

	dataMapSlice := make([]map[string]interface{}, 0)

	dataMapSlice = append(dataMapSlice, map[string]interface{}{"name": "John"})
	dataMapSlice = append(dataMapSlice, map[string]interface{}{"name": "Abel"})

	for _, data := range dataMapSlice {
		err := tmplt.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	}

	// err = tmplt.Execute(os.Stdout, data)
	// if err != nil{
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined\n",
		"notification": "{{.name}}, you've a new notification\n{{.notification}}\n",
		"error":        "Oops! An error occurred : {{.errorMessage}}\n",
	}

	parsedTemplates := make(map[string]*template.Template)
	for name, templt := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(templt))
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Printf("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var t *template.Template

		switch choice {
		case "1":
			t = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}

		case "2":
			fmt.Printf("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			t = parsedTemplates["notification"]
			data = map[string]interface{}{"name": name, "notification": notification}

		case "3":
			fmt.Printf("Enter your error message: ")
			errMessage, _ := reader.ReadString('\n')
			errMessage = strings.TrimSpace(errMessage)
			t = parsedTemplates["error"]
			data = map[string]interface{}{"name": name, "errorMessage": errMessage}

		case "4":
			fmt.Printf("Exiting...\n")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option")
			continue
		}

		err := t.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println(err)
		}
	}

}
