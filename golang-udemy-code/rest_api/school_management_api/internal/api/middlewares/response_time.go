package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
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