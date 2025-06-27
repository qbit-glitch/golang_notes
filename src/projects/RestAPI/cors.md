# CORS Middleware

CORS, which stands for Cross Origin Resource Sharing, is a security feature implemented in web-browsers that restricts web pages from making requests to a domain different from the one that served the web page. This is crucial for preventing malicious attacks, but it can be a limitation during development or when building APIs that need to be accessed from different origins. The CORS middleware allows you to configure which origins are permitted to access your resources.

Cross Origin Resourse Sharing:
- Allow Specific Origins
- HTTP Methods
- Headers
- Credentials
- Preflight Requests

When a client makes a request to a resource on a different origin for example `https://localhost:3000` to `https://api.example.com`, the browser checks whether the server's response includes the appropriate CORS headers.

All these middlewares are for production phase. In development phase we can inactivate them while we are testing our API continuously during development.

Using CORS middlware we can specify which origins are allowed to access the resources. For example we can allow requests from `localhost:3000` or from `myoriginurl.com`, etc. We are actually checking the origin header. We are accessing the origin, which needs to be as per our requirement, if it does not match our accepted origins in the origin key of the header, then in that case, we will not allow such requests. This gives us an additional security layer over your API. 

`http.MethodOptions` is just a pre-flight check. It returns immediately, allowing the request without calling the next handler.

Using CORS middleware is essential for developing APIs that need to be accessed from web-applications hosted on different origins. It helps maintaining security while allowing legitimate cross-origin requests.

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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
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