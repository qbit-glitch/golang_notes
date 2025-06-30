package handlers

import (
	"fmt"
	"net/http"
)

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
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
