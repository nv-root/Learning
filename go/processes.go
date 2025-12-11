package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func Ex1() {
	cmd := exec.Command("echo", "Hello world!")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func Ex2() {
	cmd := exec.Command("grep", "foo")
	cmd.Stdin = strings.NewReader("food bar\nbaz\n")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func PipeEx() {
	pr, pw := io.Pipe()
	cmd := exec.Command("grep", "foo")
	cmd.Stdin = pr

	go func() {
		defer pw.Close()
		pw.Write([]byte("food is good\nbar\nbaz\n"))
	}()

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func main() {
	Ex1()
	Ex2()
	PipeEx()

	cmd := exec.Command("ls", "-la")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return

	}

	fmt.Println(string(output))
}
