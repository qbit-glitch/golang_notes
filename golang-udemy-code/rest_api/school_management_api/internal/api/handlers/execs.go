package handlers

import (
	"fmt"
	"net/http"
)

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	// Find out what kind of http method that is sent with the request
	fmt.Println(r.Method)

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method Execs Route"))
	case http.MethodPost:
		fmt.Println("Query:", r.URL.Query())
		fmt.Println("Name:", r.URL.Query().Get("name"))

		// Parse form data (necessary for x-www-form-urlencoded)
		err := r.ParseForm()
		if err != nil {
			return
		}
		fmt.Println("Form from POST Method:", r.Form)

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
