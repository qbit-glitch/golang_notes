package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func studentsRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /students", handlers.GetExecsHandler)
	mux.HandleFunc("POST /students", handlers.AddExecsHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchExecsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)

	mux.HandleFunc("GET /students/{id}", handlers.GetOneExecHandler)
	mux.HandleFunc("PUT /students/{id}", handlers.UpdateStudentHandler)
	mux.HandleFunc("PATCH /students/{id}", handlers.PatchOneExecHandler)
	mux.HandleFunc("DELETE /students/{id}", handlers.DeleteOneExecHandler)

	return mux
}
