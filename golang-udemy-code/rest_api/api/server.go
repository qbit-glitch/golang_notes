package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main(){
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Handling users")
	})
	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure the TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a custom server
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})
	
	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server:", err)
	}


}