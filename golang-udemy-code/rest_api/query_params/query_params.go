package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	// Find out what kind of http method that is sent with the request
	fmt.Println(r.Method)

	switch r.Method {
	case http.MethodGet:

		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		fmt.Println("UserID:", userID)

		// teachers/?key=value&query=value2&sortby=email&sortorder=ASC

		queryParams := r.URL.Query()
		fmt.Println("Query Params:", queryParams)
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")
		if sortorder == ""{
			sortorder = "DESC"
		}
		fmt.Printf("Sortby: %v | Sort Order: %v | Key: %v", sortby, sortorder, key)

		w.Write([]byte("Hello GET method Teachers Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method Teachers Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method Teachers Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method Teachers Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method Teachers Route"))
	}

	w.Write([]byte("Hello Teachers Route"))
	fmt.Println("Hello Teachers Route")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	// Find out what kind of http method that is sent with the request
	fmt.Println(r.Method)

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method Execs Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method Execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method Execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method Execs Route"))
	}

	w.Write([]byte("Hello Execs Route"))
	fmt.Println("Hello Execs Route")
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	// Find out what kind of http method that is sent with the request
	fmt.Println(r.Method)

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method Students Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method Students Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method Students Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method Students Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method Students Route"))
	}

	w.Write([]byte("Hello Students Route"))
	fmt.Println("Hello Students Route")
}

func main() {

	port := ":3000"
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}

}
