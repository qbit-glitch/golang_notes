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