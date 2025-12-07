package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal("Error reading from reader:", err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatal("Error writing to writer:", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal("Error closing resource:", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // allocates on the stack
	buf.WriteString("Hello buffer")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader("World")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // allocates on the heap

	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatal("Error reading from multiple readers:", err)
	}

	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() { // go routine
		pw.Write([]byte("Hello io Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(pr)
	if err != nil {
		log.Fatal("Error reading from pipe reader:", err)
	}

	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening/creating file:", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("Written data to file", filepath)

	// writer := io.Writer(file)
	// writer.Write([]byte(data))
}

type MyResouce struct {
	name string
}

func (m MyResouce) Close() error {
	fmt.Println("Closing resource:", m.name)
	return nil
}

func main() {

	fmt.Println("Read from reader")
	readFromReader(strings.NewReader("I'm a string Reader, I'm a string Reader"))

	fmt.Println("\nWrite to writer")
	writeToWriter(os.Stdout, "Hello bro\n")

	fmt.Println("\nBuffer example")
	bufferExample()

	fmt.Println("\nMultireader example")
	multiReaderExample()

	fmt.Println("\nPipe example")
	pipeExample()

	writeToFile("file.txt", "Ken thompson zindabad!\n")

	resource := &MyResouce{name: "my resource"}
	closeResource(resource)

}
