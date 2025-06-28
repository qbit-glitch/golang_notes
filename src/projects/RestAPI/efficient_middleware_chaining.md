# Efficient Middleware chaining

```go
// Middleware is a function that wraps an http.Handler with additional functionality
type Middlware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
    for _, middleware := range middlewares {
        handler = middleware(handler)
    }
    return handler
}

main {
    ...
    secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
    
    server := &http.Server {
        Addr: port,
        Handler: secureMux,
        TLSConfig: tlsConfig,
    }
}
```


server.go
```go
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

	hppOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"allowedParam"},
	}

	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
    
	// create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux
		// Handler:   mw.SecurityHeaders(mux),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}

// Middleware is a function that wraps an http.Handler with additional functionality
type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
```

compression.go
```go
package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	fmt.Println("Compression Middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Compression Middleware being returned ...")
		// Check if the client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip"){
			next.ServeHTTP(w,r)
		}

		// Set the response header
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Wrap the ResponseWriter
		w = &gzipResponseWriter{
			ResponseWriter: w,
			Writer: gz,
		}

		next.ServeHTTP(w,r)
		fmt.Println("Sent from Compression Middleware")
		fmt.Println("Compression Middleware ends ...")

	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error){
	return g.Writer.Write(b)
}
```


cors.go
```go
package middlewares

import (
	"fmt"
	"net/http"
)

// api is hosted at www.myapi.com
// frontend server is at www.myfrontend.com

var allowedOrigins = []string{
	"https://localhost:3000",
	"https://www.myfrontend.com",
	"https://myoriginurl.com",
}

func Cors(next http.Handler) http.Handler {
	fmt.Println("CORS middleware starts")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("CORS middleware being returned ...")
		origin := r.Header.Get("Origin")
		fmt.Println(origin)

		if isOriginAllowed(origin){
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			http.Error(w, "Not allowed by CORS", http.StatusForbidden)
			return
		}
		// w.Header().Set()
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == http.MethodOptions{
			return
		}

		next.ServeHTTP(w, r)
		fmt.Println("CORS middleware ends ...")
	})
}

func isOriginAllowed(origin string) bool{
	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == origin {
			return true
		}
	}
	return false
}
```

hpp.go
```go
package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

type HPPOptions struct {
	CheckQuery                  bool
	CheckBody                   bool
	CheckBodyOnlyForContentType string
	Whitelist                   []string
}

func Hpp(options HPPOptions) func(http.Handler) http.Handler{
	fmt.Println("HPP middleware starts...")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("HPP middleware being returned ...")
			if options.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, options.CheckBodyOnlyForContentType){
				// Filter the body params
				filterBodyParams(r, options.Whitelist)
			}
			if options.CheckQuery && r.URL.Query() != nil {
				// Filter the query params
				filterQueryParams(r, options.Whitelist)
			}

			next.ServeHTTP(w,r)
			fmt.Println("HPP middleware ends ")
		})
	}
}

func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func filterBodyParams(r *http.Request, whitelist []string){
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	for k,v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0])    // first value
			// r.Form.Set(k, v[len(v)-1])   // last value
		}
		if !isWhiteListed(k, whitelist) {
			delete(r.Form, k)
		}
	}
}

func filterQueryParams(r *http.Request, whitelist []string){
	query := r.URL.Query()

	for k,v := range r.Form {
		if len(v) > 1 {
			query.Set(k, v[0])    // first value
			// query.Set(k, v[len(v)-1])   // last value
		}
		if !isWhiteListed(k, whitelist) {
			query.Del(k)
		}
		r.URL.RawQuery = query.Encode()
	}
}



func isWhiteListed(param string, whitelist []string) bool {
	for _, v := range whitelist {
		if param == v {
			return true
		}
	}
	return false
}
```

rate_limiter.go
```go
package middlewares

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu sync.Mutex
	visitors map[string]int
	limit int
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *rateLimiter{
	rl := &rateLimiter{
		visitors: make(map[string]int),
		limit: limit,
		resetTime: resetTime,
	}
	// Start the reset routine
	go rl.resetVisitorCount()
	return rl
}

func (rl *rateLimiter) resetVisitorCount(){
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) Middleware(next http.Handler) http.Handler{
	fmt.Println("Rate Limiter Middleware starts ...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Rate Limiter Middlerware being returned...")
		rl.mu.Lock()
		defer rl.mu.Unlock()

		visitorIP := r.RemoteAddr    // You might want to extract the IP in a more sophisticated way
		rl.visitors[visitorIP]++

		fmt.Printf("Visitor COunt from %v is %v\n", visitorIP, rl.visitors[visitorIP])

		if rl.visitors[visitorIP] > rl.limit {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w,r)
		fmt.Println("Rate Limiter Middlerware ends")
	})
}
```

response_time.go
```go
package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	fmt.Println("Response Time middleware starts ..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Response Time Middlerware being returned")
		fmt.Println("Received Request in ResponseTime")
		start := time.Now()

		// Create a custom ResponseWriter to capture the status code
		wrappedHeader := &responseWriter{
			ResponseWriter: w,
			status: http.StatusOK,
		}

		// Calculate the duration
		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())
		
		next.ServeHTTP(wrappedHeader,r)

		// Log the request details
		duration = time.Since(start)
		fmt.Printf("Method: %s, URL: %s,Status: %d, Duration: %v\n", r.Method, r.URL, wrappedHeader.status, duration.String())
		fmt.Println("Response Time Middlerware ends")
		
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int){
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
```

security.go
```go
package middlewares

import (
	"fmt"
	"net/http"
)

func SecurityHeaders(next http.Handler) http.Handler {
	fmt.Println("Security Headers Middleware starts")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Security Headers Middleware being returned")
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode-block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("X-Powered-By", "Django")

		w.Header().Set("Server", "")
		w.Header().Set("X-Permitterd-Cross-Domain-Policies", "none")
		w.Header().Set("Cache-Control", "no-store, no-chache, must-revalidate, max-age=0")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Permissions-Policy", "geolocation=(self), microphone=()")

		next.ServeHTTP(w, r)
		fmt.Println("Security Headers Middleware ends")
	})
}
```