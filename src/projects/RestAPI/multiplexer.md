# Multiplexer

In Go `mux` which is short for multiplexer, refers to a request multiplexer which is a router that matches incoming HTTP requests to their respective handlers based on request URL and method. The http.ServeMux is the default http request multiplexer provided by the Go standard libary.

Mux allows you to define multiple routes, that is multiple end points for your API. Each route can have its own handler function, enabling you to organize your API better. Also MUX helps separating the logic for different routes, making the code cleaner and more maintainable.

```go
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

	// create custom server
	server := &http.Server{
		Addr: port,
		Handler: mux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
```