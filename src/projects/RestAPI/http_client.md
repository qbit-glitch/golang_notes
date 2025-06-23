# HTTP Client

The `net/http` package provide tools to make http requests. This package allows us to create HTTP Clients that can communicate with web servers using various HTTP methods like GET, PUT, POST, PATCH and DELETE.

The core compoent for making http requests in GO is the `http.Client{}` struct. This struct can be used to send request and receive responses from a server.


## Code
```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// Create a new Client
	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/10")
	// resp, err := client.Get("https://swapi.dev/api/people/1")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	fmt.Println(body)
	fmt.Println(string(body))
}
```