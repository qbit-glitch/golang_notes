# HPP Middleware

HPP which stands for HTTP parameter pollution, is a type of an attack and our HPP middleware will help protect our application from HTTP parameter pollution attacks. Because the HTTP parameter pollution occurs when multiple parameters with the same name are sent in an HTTP request. This can lead to unexpected behavior, data corruption, or security vulnerabilities in the application. 

HPP middlware normalizes the request parameters by merging or removing duplicate parameters. It ensures that only values is kept for each parameter name, reducing the risk of ambiguity and manipulation. By preventing paramter pollution, HPP helps maintain the integrity of your applications, data and logic, making it more robust against certain types of attacks. Moreover, this middleware will also us to configure how to handle duplicate parameters such as chooding to keep the first or last occurence or merging them into an array. Overall HPP middleware adds a layer of security by ensuring that the application processes parameters in a predictable and a safe manner.


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
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if options.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, options.CheckBodyOnlyForContentType){
				// Filter the body params
				filterBodyParams(r, options.Whitelist)
			}
			if options.CheckQuery && r.URL.Query() != nil {
				// Filter the query params
				filterQueryParams(r, options.Whitelist)
			}

			next.ServeHTTP(w,r)
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