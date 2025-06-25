package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

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
		case http.MethodPost:
			w.Write([]byte("Hello POST method Teachers Route"))
			return
		case http.MethodPut:
			w.Write([]byte("Hello PUT method Teachers Route"))
			return
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH method Teachers Route"))
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
