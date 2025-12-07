package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed test
var testFolder embed.FS

func main() {

	fmt.Println("Embedded content:", content)
	content, err := testFolder.ReadFile("test/hello.txt")
	if err != nil {
		fmt.Println("Error reading file:\n", err)
		return
	}

	fmt.Println("Embedded file content:\n", string(content))

	err = fs.WalkDir(testFolder, "test", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println("PATH -", path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
