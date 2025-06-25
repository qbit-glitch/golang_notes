# Basic Routing - CRUD - HTTP Methods


**CRUD OPERATIONS**
- Create
- Read
- Update
- Delete

**HTTP Methods**
- POST
- GET
- PUT
- DELETE
- PATCH


The http methods are fundamental to the design and operation of web applications and APIs. They provide a standardized way to perform different actions on resources, making web services predictable, interoperable and easy to understand. Moreover the RESTful API design REST, which is Representational State Transfer, is an architectural fo designing networked applications and RESTful APIs use HTTP methods to perform CRUD operations on resources. This approach simplifies the design and makes APIs more intuitive.

To make the APIs simple and intuitive, all the methods that are related to a particular are grouped together under one root.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := ":3000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		// Find out what kind of http method that is sent with the request
		fmt.Println(r.Method)

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET method Teachers Route"))
			return
		case http.MethodPut:
			w.Write([]byte("Hello PUT method Teachers Route"))
			return
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH method Teachers Route"))
			return
		case http.MethodPost:
			w.Write([]byte("Hello POST method Teachers Route"))
			return
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE method Teachers Route"))
			return
		}

		w.Write([]byte("Hello Teachers Route"))
		fmt.Println("Hello Teachers Route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Execs Route"))
		fmt.Println("Hello Execs Route")
	})

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
```