package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.RootHandler)

	mux.HandleFunc("GET /teachers", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /teachers", handlers.AddStudentsHandler)
	mux.HandleFunc("PATCH /teachers", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /teachers", handlers.DeleteStudentsHandler)

	mux.HandleFunc("GET /teachers/{id}", handlers.GetOneStudentHandler)
	mux.HandleFunc("PUT /teachers/{id}", handlers.UpdateStudentHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchOneStudentHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteOneStudentHandler)

	mux.HandleFunc("GET /students", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students", handlers.AddStudentsHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)

	mux.HandleFunc("GET /students/{id}", handlers.GetOneStudentHandler)
	mux.HandleFunc("PUT /students/{id}", handlers.UpdateStudentHandler)
	mux.HandleFunc("PATCH /students/{id}", handlers.PatchOneStudentHandler)
	mux.HandleFunc("DELETE /students/{id}", handlers.DeleteOneStudentHandler)

	mux.HandleFunc("/execs/", handlers.ExecsHandler)

	return mux
}
