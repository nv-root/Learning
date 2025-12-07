package main

import "os"
import "fmt"
import "bufio"
import "strings"

func main() {

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
defer file.Close()

	scanner := bufio.NewScanner(file)
	keyword := "Hello"

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {

			updatedLine := strings.ReplaceAll(line, keyword, "Hi")
			fmt.Println("Filtered line:", line)
			fmt.Println("Updated line:", updatedLine)
		}

		err = scanner.Err()
		if err != nil {
			fmt.Println("Error scanning file", err)
			file.Close()
			return
		}

	}

}
