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