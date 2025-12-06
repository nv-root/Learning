// base64 - bin -> text using 64 ascii chars
package main

import "fmt"
import "encoding/base64"
func main(){

	data := []byte("Hello, base64 encoding!@#$%^&*()?/`~`")

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("base64 encoded string:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil{
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println(decoded)
	fmt.Println("decoded string:", string(decoded))

	// url safe encoding removes +, /
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("\nURL safe encoded:",urlSafeEncoded)
	decoded, err = base64.URLEncoding.DecodeString(urlSafeEncoded)
	fmt.Println(decoded)
	fmt.Println(string(decoded))

	fmt.Println()
	fmt.Println(encoded)
	fmt.Println(urlSafeEncoded)
}
