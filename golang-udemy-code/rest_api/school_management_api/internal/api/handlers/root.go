package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to our School API"))
	// fmt.Println("Hello Root Route")
}
