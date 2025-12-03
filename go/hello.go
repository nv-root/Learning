package main

import (
	format "fmt"
	"net/http"
)

func main() {
	format.Println("Hello Go!")

	res, err := http.Get("https://api.github.com/users/nv-root")
	if err != nil {
		format.Println("Error: ", err)
		return
	}

	defer res.Body.Close()
	format.Println(res.Status)
}
