package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	
	mux.HandleFunc("GET /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("GET /teachers/{id}", handlers.TeachersHandler)
	mux.HandleFunc("POST /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PUT /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PATCH /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.TeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.TeachersHandler)
	
	
	
	mux.HandleFunc("/students/", handlers.StudentsHandler)
	mux.HandleFunc("/execs/", handlers.ExecsHandler)
	return mux
}
