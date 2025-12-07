package main

import "fmt"
import "strconv"
import "strings"
import "regexp"

func main() {
	str := "Hello Go!"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"

	result := str1 + " " + str2

	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[:5])

	num := 65
	str = strconv.Itoa(num)
	fmt.Println(len(str)) // 65 -> "65"

	fruits_str := "apple,banana,cherry"
	fruits := strings.Split(fruits_str, ",")
	fmt.Println(fruits)

	fruits_str = "apple-banana-cherry"
	fruits = strings.Split(fruits_str, "-")
	fmt.Println(fruits)

	fruits_str = strings.Join(fruits, "::")
	fmt.Println(fruits_str)

	//  checking for substring -> bool
	fmt.Println(strings.Contains("Hello Go!", "llo"))

	// replacing a substring
	replaced := strings.Replace("Hello Go! Go!", "Go", "World", 2)
	fmt.Println(replaced)

	fmt.Println(strings.TrimSpace(" Hello   "))

	fmt.Println(strings.ToLower("Hello"))
	fmt.Println(strings.ToUpper("Hello"))

	fmt.Println(strings.Repeat("foo", 3))

	fmt.Println(strings.Count("foo", "f"))

	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasSuffix("Hello", "-"))

	re := regexp.MustCompile(`\d+`)                    // one or more digits
	matches := re.FindAllString("Hello 1 111 123", -1) // -1 - all

	fmt.Println(re)
	fmt.Println(matches)


	// string building
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	builderResult := builder.String()
	fmt.Println(builderResult)

	builder.WriteRune('!')
	builder.WriteString("How are you?")

	builderResult = builder.String()
	fmt.Println(builderResult)

	// resetting builder
	builder.Reset()
	builder.WriteString("How are you?")
	builderResult = builder.String()
	fmt.Println(builderResult)
		

}
