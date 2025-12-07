package main

import "os"
import "fmt"
import "path/filepath"

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		panic(err)
	}
}

func main() {

	// err := os.Mkdir("subdir", 0755)
	err := os.MkdirAll("subdir/sub-subdir/", 0755) // creates subdirs if not exists and remains silent if already exists unlike Mkdir
	checkError("Error creating dir:", err)

	os.WriteFile("subdir/file.txt", []byte("yo!"), 0755)
	os.WriteFile("subdir/file2.txt", []byte("yo!"), 0755)

	//reading dir contents
	result, err := os.ReadDir(".")
	checkError("Error reading dir:", err)

	for _, entry := range result {
		fmt.Println(entry.Name() /*, entry.IsDir(), entry.Type()*/)
	}

	// changin dir
	fmt.Println("\nGoing to child dir...")
	err = os.Chdir("./subdir/")
	checkError("Error changing dir:", err)

	result, err = os.ReadDir(".")
	checkError("Error reading dir:", err)

	for _, entry := range result {
		fmt.Println(entry.Name() /*, entry.IsDir(), entry.Type()*/)
	}

	// getting working dir
	dir, err := os.Getwd()
	checkError("Error getting working dir:", err)
	fmt.Println("Current dir:", dir)

	err = os.Chdir("../")
	checkError("Error changing dir:", err)

	fmt.Println("\nGoing back to parent...")
	dir, err = os.Getwd()
	checkError("Error getting working dir:", err)
	fmt.Println("Current dir:", dir)

	// ==========================
	// using filepath
	// ==========================
	pathfile := "."
	fmt.Println("\n\nWalking dir...")
	filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error : ", err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	defer os.RemoveAll("subdir")

}
