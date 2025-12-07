package main

import "fmt"
import "path/filepath"
import "strings"

func main() {
	relativePath := "../README.md"
	absolutePath := "/home/vamsi/Programming/Learning/README.md"

	joinedPath := filepath.Join("downloads","main.c")
	fmt.Println("joined path:", joinedPath)

	dir, file := filepath.Split(absolutePath)
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)

	base := filepath.Base(absolutePath)
	fmt.Println("base:", base)

	fmt.Println(filepath.IsAbs(absolutePath))
	fmt.Println(filepath.IsAbs(relativePath))

	fmt.Println("extension:", filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file))) // removing extension

	rel, err := filepath.Rel("a/c", "a/b/c/file.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Abs path:", absPath)
	}



}
