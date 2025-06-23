# HTTP Server

The `net/http` package provides robust tools for building http servers. Understanding how to create and manage HTTP servers is fundamental for developing web applications and APIs. To create an HTTP server in Go, we need to define HTTP Handlers. HTTP handlers are functions that handle the logic, the business logic for an endpoint. User sends request to a specific endpoint, then the handler will be executed when that endpoint receives a request from a client. And apart from defining http handlers, we also configure routes and we need to configure that in our API in our server. And finally we start the server to listen on a specific port.

## Code
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "Hello World")
	})
	// const serverAddr string = "127.0.0.1:8080"
	const port string = ":8080"
	fmt.Println("Server Listening on Port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
```