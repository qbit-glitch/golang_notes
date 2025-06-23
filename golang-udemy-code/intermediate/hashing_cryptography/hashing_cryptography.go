package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	// simpleHashingExamples()

	passwordHashingExample()

}

func simpleHashingExamples(){
	password := "password123"

	hash256 := sha256.Sum256([]byte (password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println("SHA 256 Hash:", hash256)
	fmt.Println("SHA 512 Hash:", hash512)

	fmt.Printf("SHA 256 hex value: %x\n", hash256)
	fmt.Printf("SHA 512 hex value: %x\n", hash512)
}

func passwordHashingExample(){

	password := "password123"

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt: %x\n", salt)

	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and password in database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)  // Simulate as storing in database
	fmt.Println("Hash:", signUpHash)  // SImulate as storing in database


	// Verify
	// Retrieve the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("unable to decode salt:", err)
		return
	}

	passwordToBeChecked := password + "1234"
	loginHash := hashPassword(passwordToBeChecked, decodedSalt)

	// compare the stored signup hash with LoginChain
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login Failed.Please check user credentials.")
	}




}

func generateSalt() ([]byte, error){
	salt := make([]byte, 20)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte (password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])

}