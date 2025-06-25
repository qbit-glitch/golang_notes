package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age string `json:"age"`
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

			fmt.Println("Body:", r.Body)
			fmt.Println("Form:", r.Form)
			fmt.Println("Header:", r.Header)
			fmt.Println("Context:", r.Context())
			fmt.Println("Content Length:", r.ContentLength)
			fmt.Println("Host:", r.Host)
			fmt.Println("Method:", r.Method)
			fmt.Println("Protocol:", r.Proto)
			fmt.Println("Remote Addr:", r.RemoteAddr)
			fmt.Println("Request URI:", r.RequestURI)
			fmt.Println("TLS:", r.TLS)
			fmt.Println("Trailer:", r.Trailer)
			fmt.Println("Transfer Encoding:", r.TransferEncoding)
			fmt.Println("URL:", r.URL)
			fmt.Println("User Agent", r.UserAgent())
			fmt.Println("Port:", r.URL.Port())
			
			w.Write([]byte("Hello GET method Teachers Route"))
			return
		case http.MethodPost:
			// Parse form data (necessary for x-www-form-urlencoded)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form:", http.StatusBadRequest)
				return
			}
			fmt.Println("Form:", r.Form)

			// Parse Response Data
			response := make(map[string]interface{})
			for key, value := range r.Form {
				response[key] = value[0]
			}
			fmt.Println("Processed Response Map:", response)

			// RAW Body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error parsing form:", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			fmt.Println("Raw Body:", string(body))


			// If you expect JSON data, then Unmarshall it using struct
			var userInstance1 User
			err = json.Unmarshal(body, &userInstance1)
			if err != nil {
				http.Error(w, "Error parsing form:", http.StatusBadRequest)
				return
			}
			fmt.Println(userInstance1)

			// Using maps to Unmarshal the data
			userInstance2 := make(map[string]interface{})
			err = json.Unmarshal(body, &userInstance2)
			if err != nil {
				http.Error(w, "Error parsing form:", http.StatusBadRequest)
				return
			}
			fmt.Println("Unmarshaled JSON into a map", userInstance2)

			fmt.Println("Body:", r.Body)
			fmt.Println("Form:", r.Form)
			fmt.Println("Header:", r.Header)
			fmt.Println("Context:", r.Context())
			fmt.Println("Content Length:", r.ContentLength)
			fmt.Println("Host:", r.Host)
			fmt.Println("Method:", r.Method)
			fmt.Println("Protocol:", r.Proto)
			fmt.Println("Remote Addr:", r.RemoteAddr)
			fmt.Println("Request URI:", r.RequestURI)
			fmt.Println("TLS:", r.TLS)
			fmt.Println("Trailer:", r.Trailer)
			fmt.Println("Transfer Encoding:", r.TransferEncoding)
			fmt.Println("URL:", r.URL)
			fmt.Println("User Agent", r.UserAgent())
			fmt.Println("Port:", r.URL.Port())
			

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
