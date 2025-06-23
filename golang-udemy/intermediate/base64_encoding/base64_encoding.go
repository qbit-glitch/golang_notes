package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Base64 encoding")
	fmt.Println(data)

	// Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded String:", encoded)

	// DEcode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("error in decoding:", err)
		return
	}
	fmt.Println("Decoded Value:",decoded)
	fmt.Println("Decoded String:", string(decoded))

	// URL Safe, avoid '/' and '+'
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe ENcoding:", urlSafeEncoded)

	
}