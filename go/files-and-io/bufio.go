package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// accepts another reader and buffers it
	reader := bufio.NewReader(strings.NewReader("Hello, bufio package woah\n"))

	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading bytes:", err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(data))

	line, _ := reader.ReadString('\n')
	fmt.Println(line)

	// writing
	writer := bufio.NewWriter(os.Stdout)
	dataToWrite := []byte("Hello, bufioooooo package\n")
	n, err = writer.Write(dataToWrite)

	if err != nil {
		fmt.Println("Error writing bytes:", err)
	}

	fmt.Println(n, "bytes written to stdout")

	_ = writer.Flush()

	str := "This is a string\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}

	fmt.Println(n, "bytes written to stdout")
	_ = writer.Flush()

}
