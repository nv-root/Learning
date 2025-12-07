package main

import (
	"fmt"
	"os"
)

func main() {

	// creates temp files and dirs in the default  temporary dir of the os if not the provided the directory arg like done here
	tempFile, err := os.CreateTemp("", "temporaryFile")
	if err != nil {
		panic(err)

	}

	fmt.Println("Temp file created:", tempFile.Name())

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	tempDir, err := os.MkdirTemp("", "temporaryDirectory")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempDir)
	fmt.Println("Temp dir created:", tempDir)

}
