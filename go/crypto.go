// hashing is irreversable, deterministic
package main

import "fmt"
import "crypto/sha256"
import "crypto/sha512"
import "crypto/rand"
import "encoding/base64"
import "io"


func genSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)

	return base64.StdEncoding.EncodeToString(hash[:])
}

func main() {

	var password string
	fmt.Printf("Enter signup password: ")
	fmt.Scan(&password)

	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println("Entered password:", password)

	// fmt.Println("hash 256:", hash256)
	fmt.Printf("hash256 hex: %x\n", hash256)

	// fmt.Println("hash 512:", hash512)
	fmt.Printf("hash512 hex: %x\n", hash512)


	// generating the salt
	salt, err := genSalt()
	if err != nil{
		fmt.Println("Error generating salt:", err)
		return
	}

	hash := hashPassword(password, salt)

	// base64 encoding the salt to store in db
	saltStr := base64.StdEncoding.EncodeToString(salt)

	fmt.Printf("orignal salt hex: %x\n", salt)
	fmt.Println("encoded salt:", saltStr)
	fmt.Println("salted and hashed password:", hash)


	// ==========================
	// verifying
	// ==========================

	fmt.Printf("Enter login password: ")
	fmt.Scan(&password)
	decodedSalt, _ := base64.StdEncoding.DecodeString(saltStr)
	loginHash := hashPassword(password, decodedSalt)

	// comparing
	if hash == loginHash{
		fmt.Println("Password is correct")
	}else{
		fmt.Println("Invalid credentials")
	}

}
