package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	
	mux.HandleFunc("GET /teachers/", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers/", handlers.AddTeacherHandler)
	mux.HandleFunc("PATCH /teachers/", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.DeleteTeachersHandler)
	
	mux.HandleFunc("GET /teachers/{id}", handlers.GetOneTeacherHandler)
	mux.HandleFunc("PUT /teachers/{id}", handlers.UpdateTeacherHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchOneTeacherHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteTeachersHandler)
	
	
	mux.HandleFunc("/students/", handlers.StudentsHandler)
	mux.HandleFunc("/execs/", handlers.ExecsHandler)
	return mux
}
