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
