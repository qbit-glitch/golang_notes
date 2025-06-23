# Simple API server

## Code
```go
package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handling users")
	})
	port := 8080
	fmt.Println("Server is running on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
```