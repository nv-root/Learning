// there are a bunch of bad practices here eg.,reusing file and err variables
// always close file after an operation
package main

import "os"
import "fmt"
import "bufio"

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		file.Close()
		return
	}

	// writing bytes to file
	data := []byte("Hello world\n\n\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		file.Close()
		return
	}

	fmt.Println("Data written to output.txt")

	// writing a string to file
	_, err = file.WriteString("Hello go")
	if err != nil {
		fmt.Println("Error writing string to output.txt")
		file.Close()
		return
	}

	fmt.Println("Data string written to output.txt")

	file, err = os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		file.Close()
		return
	}

	fmt.Println("File opened\n")

	// reading contents of the file
	data = make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		file.Close()
		return
	}

	fmt.Println("file content:")
	fmt.Println(string(data))



	_, err = file.Seek(0, 0) // resetting file pointer to the start of the file
	// reading line-by-line
	scanner := bufio.NewScanner(file)
	i := 1
	fmt.Println("\n--------------------------------------")
	for scanner.Scan() {
		line := scanner.Text() // returns the current line
		fmt.Printf("%d.| %s\n", i, line)
		i++
	}
	fmt.Println("--------------------------------------")
	err = scanner.Err()
	if err != nil{
		fmt.Println("Error reading file:", err)
		file.Close()
		return
	}



	defer file.Close()
}
