package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "school_management_api/internal/api/middlewares"
	"time"
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
	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler) 
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions {
		CheckQuery: true,
		CheckBody: true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist: []string{"allowedParam"},
	}

	secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))

	// create custom server
	server := &http.Server{
		Addr:      port,
		// Handler: mux
		// Handler:   mw.SecurityHeaders(mux),
		Handler: secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
