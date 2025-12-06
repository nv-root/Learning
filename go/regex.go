package main

import "fmt"
import "regexp"

func main() {

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "invalid_email"

	fmt.Println("Email:", re.MatchString(email1))
	fmt.Println("Email:", re.MatchString(email2))



	// capturing groups
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	data := "2025-12-05"

	submatches := re.FindStringSubmatch(data)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])
	
	re = regexp.MustCompile(`[aeiou]`)
	str := "Hello world"
	
	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)


	re = regexp.MustCompile(`(?i)go`) // ?i - case-insensitive flag

	str = "Golang is going great"
	fmt.Println(re.MatchString(str))

}
