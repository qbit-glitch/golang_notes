package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"golang.org/x/net/http2"
)

// Person struct
type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// Sample Data
var personData = map[string]Person{
	"1": {Name: "John Doe", Age: 30},
	"2": {Name: "Jane Doe", Age: 23},
	"3": {Name: "Jack Doe", Age: 25},
}

// Handler function for the endpoint
func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}
	// Check if the ID exists in the personData map
	person, exists := personData[id]
	if !exists {
		http.Error(w, "Person Nor found", http.StatusNotFound)
		return
	}
	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the person data to JSON and write to the response
	if err := json.NewEncoder(w).Encode(person); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {

	// define the port
	port := 8000

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure the TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a custom server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		TLSConfig:    tlsConfig,
		// TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)), // make it http1.1 with tls
	}

	// Set up the endpoint and the handler function
	http.HandleFunc("/person", getPersonHandler)

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server:", err)
	}

}
