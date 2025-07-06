# Part - 4

## Student Database Creation

```sql
USE school;
CREATE INDEX idx_class ON teachers(class);

CREATE TABLE IF NOT EXISTS students (
	id INT AUTO_INCREMENT PRIMARY KEY,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	class VARCHAR(255) NOT NULL,
	INDEX(email),
	FOREIGN KEY (class) REFERENCES teachers(class)
) AUTO_INCREMENT=1000
```

## CRUD for Students Route
`internal/api/handlers/students.go`
```go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"school_management_api/internal/models"
	"school_management_api/internal/repository/sqlconnect"
	"strconv"
	// "sync"
)

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var students []models.Student
	students, err := sqlconnect.GetStudentsDbHandler(students, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Student `json:"data"`
	}{
		Status: "success",
		Count:  len(students),
		Data:   students,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	// Handle Path parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// fmt.Println(err)
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	student, err := sqlconnect.GetStudentByID(id)
	if err != nil {
		// fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func AddStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var newStudents []models.Student
	var rawStudents []map[string]interface{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request Body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &newStudents)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	fmt.Println(rawStudents)

	fields := GetFieldNames(models.Student{})

	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	for _, student := range rawStudents {
		for key := range student {
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable field found in request. Only use allowed fields.", http.StatusBadRequest)
			}
		}
	}

	err = json.Unmarshal(body, &rawStudents)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		fmt.Println("New Students:", newStudents)
		return
	}

	for _, student := range newStudents {
		err = CheckBlankFields(student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedStudents, err := sqlconnect.AddStudentsDBHandler(newStudents)
	if err != nil {
		// fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Student `json:"data"`
	}{
		Status: "success",
		Count:  len(addedStudents),
		Data:   addedStudents,
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var updatedStudent models.Student
	err = json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedStudentFromDb, err := sqlconnect.UpdateStudent(id, updatedStudent)
	if err != nil {
		// log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedStudentFromDb)

}

// PATCH /students/
func PatchStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchStudents(updates)
	if err != nil {
		// log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// PATCH /students/{id}
func PatchOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedStudent, err := sqlconnect.PatchOneStudent(id, updates)
	if err != nil {
		// log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedStudent)

}

func DeleteOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneStudent(id)
	if err != nil {
		// log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Response Body -> Optional
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Student deleted successfully",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)

	// Return status of NoContent -> Compulsory
	w.WriteHeader(http.StatusNoContent)

}

func DeleteStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var ids []int
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	deleteIds, err := sqlconnect.DeleteStudents(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status     string `json:"status"`
		DeletedIDs []int  `json:"deleted_ids"`
	}{
		Status:     "Students successfully deleted",
		DeletedIDs: deleteIds,
	}
	json.NewEncoder(w).Encode(response)
}
```

`repository/sqlconnect/students_crud.go`
```go
package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"school_management_api/internal/models"
	"school_management_api/pkg/utils"
	"strconv"
	"strings"
)

func GetStudentsDbHandler(students []models.Student, r *http.Request) ([]models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, class FROM students WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("err")
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer rows.Close()
	for rows.Next() {
		student := models.Student{}
		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {

			return nil, utils.ErrorHandler(err, "error retrieving data")
		}
		students = append(students, student)
	}
	return students, nil
}

func GetStudentByID(id int) (models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error retrieving data ")
	}
	defer db.Close()

	var student models.Student
	err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
	if err == sql.ErrNoRows {
		return models.Student{}, utils.ErrorHandler(err, "error retrieving data ")
	} else if err != nil {
		fmt.Println(err)
		return models.Student{}, utils.ErrorHandler(err, "error retrieving data ")
	}
	return student, nil
}

func AddStudentsDBHandler(newStudents []models.Student) ([]models.Student, error) {

	fmt.Println("------ AddStudentsDBHandler Called -------")

	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error adding data")
	}
	defer db.Close()
	stmt, err := db.Prepare(utils.GenerateInsertQuery("students", models.Student{}))
	if err != nil {
		return nil, utils.ErrorHandler(err, "error adding data")
	}
	defer stmt.Close()
	fmt.Printf("Studentss Add Handler")

	addedStudents := make([]models.Student, len(newStudents))

	for i, newStudent := range newStudents {
		values := utils.GetStructValues(newStudent)
		fmt.Println(newStudent)

		fmt.Println("VALUES:", values)
		res, err := stmt.Exec(values...)

		if err != nil {
			fmt.Println("----- Error():", err)
			if strings.Contains(err.Error(), "a foreign key constraint fails (`school`.`students`, CONSTRAINT `students_ibfk_1` FOREIGN KEY (`class`) REFERENCES `teachers` (`class`))"){
				return nil, utils.ErrorHandler(err, "class / class teacher does not exist.")
			}
			return nil, utils.ErrorHandler(err, "error adding data")
		}
		lastID, err := res.LastInsertId()
		if err != nil {

			return nil, utils.ErrorHandler(err, "error adding data")
		}
		newStudent.ID = int(lastID)
		addedStudents[i] = newStudent
	}
	return addedStudents, nil
}

func UpdateStudent(id int, updatedStudent models.Student) (models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, class, email, first_name, last_name FROM students WHERE id = ?", id).Scan(&existingStudent.ID, &existingStudent.Class, &existingStudent.Email, &existingStudent.FirstName, &existingStudent.LastName)
	if err != nil {
		if err != sql.ErrNoRows {

			return models.Student{}, utils.ErrorHandler(err, "error updating data")
		}
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}

	updatedStudent.ID = existingStudent.ID
	_, err = db.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", updatedStudent.FirstName, updatedStudent.LastName, updatedStudent.Email, updatedStudent.Class, updatedStudent.ID)
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	return updatedStudent, nil
}

func PatchStudents(updates []map[string]interface{}) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()

			return utils.ErrorHandler(err, "Invalid id")
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "invalid id")
		}

		var studentFromDb models.Student
		err = db.QueryRow("SELECT id, first_name, last_name, email, class FROM students WHERE id = ?", id).Scan(&studentFromDb.ID, &studentFromDb.FirstName, &studentFromDb.LastName, &studentFromDb.Email, &studentFromDb.Class)

		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {

				return utils.ErrorHandler(err, "Student Not Found")
			}

			return utils.ErrorHandler(err, "error updating data")
		}

		studentVal := reflect.ValueOf(&studentFromDb).Elem()
		studentType := studentVal.Type()

		for k, v := range update {
			if k == "id" {
				continue
			}
			for i := 0; i < studentVal.NumField(); i++ {
				field := studentType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := studentVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("cannot convert %v to %v", val.Type(), fieldVal.Type())
							return utils.ErrorHandler(err, "error updating data")
						}
					}
					break
				}
			}
		}
		_, err = tx.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", studentFromDb.FirstName, studentFromDb.LastName, studentFromDb.Email, studentFromDb.Class, studentFromDb.ID)
		if err != nil {
			tx.Rollback()

			return utils.ErrorHandler(err, "error updating data")
		}
	}
	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}
	return nil
}

func PatchOneStudent(id int, updates map[string]interface{}) (models.Student, error) {

	db, err := ConnectDb()
	if err != nil {
		log.Println(err)
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	var existingStudent models.Student
	err = db.QueryRow("SELECT id, class, email, first_name, last_name FROM students WHERE id = ?", id).Scan(&existingStudent.ID, &existingStudent.Class, &existingStudent.Email, &existingStudent.FirstName, &existingStudent.LastName)
	if err != nil {
		if err != sql.ErrNoRows {

			return models.Student{}, utils.ErrorHandler(err, "Student not Found")
		}
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	studentVal := reflect.ValueOf(&existingStudent).Elem()
	studentType := studentVal.Type()
	for k, v := range updates {

		for i := 0; i < studentVal.NumField(); i++ {

			field := studentType.Field(i)

			if field.Tag.Get("json") == k+",omitempty" {
				if studentVal.Field(i).CanSet() {
					studentVal.Field(i).Set(reflect.ValueOf(v).Convert(studentVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, class = ? WHERE id = ?", existingStudent.FirstName, existingStudent.LastName, existingStudent.Email, existingStudent.Class, existingStudent.ID)
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error updating data")
	}
	return existingStudent, nil
}

func DeleteOneStudent(id int) error {
	db, err := ConnectDb()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "error deleting data")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return utils.ErrorHandler(err, "error deleting data")
	}

	fmt.Println(result.RowsAffected())

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "error deleting data")
	}

	if rowsAffected == 0 {
		return utils.ErrorHandler(err, "student not found")
	}
	return nil
}

func DeleteStudents(ids []int) ([]int, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error deleting data")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error deleting data")
	}

	stmt, err := tx.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return nil, utils.ErrorHandler(err, "error deleting data")
	}
	defer stmt.Close()

	deleteIds := []int{}

	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()

			return nil, utils.ErrorHandler(err, "error deleting data")
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()

			return nil, utils.ErrorHandler(err, "error deleting data")
		}

		if rowsAffected > 0 {
			deleteIds = append(deleteIds, id)
		}

		if rowsAffected < 1 {
			tx.Rollback()

			return nil, utils.ErrorHandler(err, fmt.Sprintf("ID %d does not exist", id))
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrorHandler(err, "error deleting data")
	}
  
	if len(deleteIds) < 1 {
		return nil, utils.ErrorHandler(err, "IDs do not exist")
	}
	return deleteIds, nil
}
```

## New Subroutes

Never use `nil` as a handler, otherwise you would always get an error and your server won't start. Subroutes are sub-urls. Any router after the main URL is the subroute.

`router/router.go`
```go
mux.HandleFunc("GET /teachers/{id}/students", handlers.GetStudentsByTeacherId)
mux.HandleFunc("GET /teachers/{id}/studentcount", handlers.GetStudentCountByTeacherId)
```

## Getting Student List for a specific teacher

`teachers.go`
```go
func GetStudentsByTeacherId(w http.ResponseWriter, r *http.Request) {
	teacherId := r.PathValue("id")

	var students []models.Student

	students, err := sqlconnect.GetStudentsByTeacherIdFromDb(teacherId, students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Student `json:"data"`
	}{
		Status: "success",
		Count:  len(students),
		Data:   students,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
```

`teachers_crud.go`
```go

func GetStudentsByTeacherIdFromDb(teacherId string, students []models.Student) ([]models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	query := `SELECT id, first_name, last_name, email, class FROM students WHERE class = (SELECT class FROM teachers WHERE id = ?)`
	rows, err := db.Query(query, teacherId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			return nil, utils.ErrorHandler(err, "error retrieving data")
		}
		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	return students, nil
}
```

## Getting Student Count for a specific teacher

Well, the student list can be a lot longer than we think, and it may take time to be generated and sometimes it may be possible that the client only needs the count and not the list. That's why a separate handler for counting the students.

`teachers.go`
```go
func GetStudentCountByTeacherId(w http.ResponseWriter, r *http.Request) {
	teacherId := r.PathValue("id")
	var studentCount int

	studentCount, err := sqlconnect.GetStudentCountByTeacherIdFromDb(teacherId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string `json:"status"`
		Count  int    `json:"count"`
	}{
		Status: "success",
		Count:  studentCount,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
```


`teachers_crud.go`
```go
func GetStudentCountByTeacherIdFromDb(teacherId string) (int, error) {
	db, err := ConnectDb()
	if err != nil {
		return 0, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	query := `SELECT COUNT(*) FROM students WHERE class = (SELECT class FROM teachers WHERE id = ?)`
	var studentCount int
	err = db.QueryRow(query, teacherId).Scan(&studentCount)
	if err != nil {
		return 0, utils.ErrorHandler(err, "error retrieving data")
	}
	return studentCount, nil
}
```

## Router Refactoring

`router/router.go`
```go
package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {

	tRouter := teachersRouter()
	sRouter := studentsRouter()

	tRouter.Handle("/", sRouter)
	return tRouter
}
```

`router/students_router.go`
```go
package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func studentsRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /students", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students", handlers.AddStudentsHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)

	mux.HandleFunc("GET /students/{id}", handlers.GetOneStudentHandler)
	mux.HandleFunc("PUT /students/{id}", handlers.UpdateStudentHandler)
	mux.HandleFunc("PATCH /students/{id}", handlers.PatchOneStudentHandler)
	mux.HandleFunc("DELETE /students/{id}", handlers.DeleteOneStudentHandler)

	return mux
}
```

`router/teachers_router.go`
```go
package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func teachersRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.RootHandler)

	mux.HandleFunc("GET /teachers", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers", handlers.AddTeachersHandler)
	mux.HandleFunc("PATCH /teachers", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers", handlers.DeleteTeachersHandler)

	mux.HandleFunc("GET /teachers/{id}", handlers.GetOneTeacherHandler)
	mux.HandleFunc("PUT /teachers/{id}", handlers.UpdateTeacherHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchOneTeacherHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteOneTeacherHandler)

	mux.HandleFunc("GET /teachers/{id}/students", handlers.GetStudentsByTeacherId)
	mux.HandleFunc("GET /teachers/{id}/studentcount", handlers.GetStudentCountByTeacherId)

	return mux
}
```

## Execs Router

`router/execs_router.go`
```go
package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func execsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /execs", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs", handlers.ExecsHandler)
	mux.HandleFunc("PATCH /execs", handlers.ExecsHandler)

	mux.HandleFunc("GET /execs/{id}", handlers.ExecsHandler)
	mux.HandleFunc("PATCH /execs/{id}", handlers.ExecsHandler)
	mux.HandleFunc("DELETE /execs/{id}", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/{id}/updatepassword", handlers.ExecsHandler)
	
	mux.HandleFunc("POST /execs/login", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/logout", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/forgotpassword", handlers.ExecsHandler)
	mux.HandleFunc("POST /execs/resetpassword/reset/{resetcode}", handlers.ExecsHandler)
	return mux
}
```

`router/router.go`
```go
package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {

	eRouter := execsRouter()
	tRouter := teachersRouter()
	sRouter := studentsRouter()
	sRouter.Handle("/", eRouter)
	tRouter.Handle("/", sRouter)
	return tRouter
}

```


## Execs Models and Database Table

Create the execs database using this query:
```sql
CREATE TABLE IF NOT EXISTS execs (
	id INT AUTO_INCREMENT PRIMARY KEY,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	username VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	password_changed_at VARCHAR(255),
	user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	password_reset_token VARCHAR(255),
	inactive_status BOOLEAN NOT NULL,
	role VARCHAR(50) NOT NULL,
	INDEX idx_email (email),
	INDEX idx_username (username)
);
```

`models/execs.go`
```go
package models

import "database/sql"

type Exec struct {
	ID                  int
	FirstName           string
	LastName            string
	Email               string
	Username            string
	Password            string
	PasswordChangedAt   sql.NullString
	UserCreatedAt       sql.NullString
	PasswordResetCode   sql.NullString
	PasswordCodeExpires sql.NullString
	InactiveStatus      bool
	Role                string
}
```

## CRUD for Execs Route

`internal/api/handlers/execs.go`
```go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"school_management_api/internal/models"
	"school_management_api/internal/repository/sqlconnect"
	"strconv"
)

func GetExecsHandler(w http.ResponseWriter, r *http.Request) {

	var execs []models.Exec
	execs, err := sqlconnect.GetExecsDbHandler(execs, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(execs),
		Data:   execs,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	// Handle Path parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// fmt.Println(err)
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	exec, err := sqlconnect.GetExecByID(id)
	if err != nil {
		// fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exec)
}

func AddExecsHandler(w http.ResponseWriter, r *http.Request) {

	var newExecs []models.Exec
	var rawExecs []map[string]interface{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request Body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &newExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	fmt.Println(rawExecs)

	fields := GetFieldNames(models.Exec{})

	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	for _, exec := range rawExecs {
		for key := range exec {
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable field found in request. Only use allowed fields.", http.StatusBadRequest)
			}
		}
	}

	err = json.Unmarshal(body, &rawExecs)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		fmt.Println("New Execs:", newExecs)
		return
	}

	for _, exec := range newExecs {
		err = CheckBlankFields(exec)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedExecs, err := sqlconnect.AddExecsDBHandler(newExecs)
	if err != nil {
		// fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(addedExecs),
		Data:   addedExecs,
	}

	json.NewEncoder(w).Encode(response)
}

// PATCH /execs/
func PatchExecsHandler(w http.ResponseWriter, r *http.Request) {

	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchExecs(updates)
	if err != nil {
		// log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// PATCH /execs/{id}
func PatchOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Exec ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedExec, err := sqlconnect.PatchOneExec(id, updates)
	if err != nil {
		// log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedExec)

}

func DeleteOneExecHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Exec ID", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneExec(id)
	if err != nil {
		// log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Response Body -> Optional
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Exec deleted successfully",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)

	// Return status of NoContent -> Compulsory
	w.WriteHeader(http.StatusNoContent)

}
```

`internal/repository/sqlconnect/execs_crud.go`
```go
package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"school_management_api/internal/models"
	"school_management_api/pkg/utils"
	"strconv"
)

func GetExecsDbHandler(execs []models.Exec, r *http.Request) ([]models.Exec, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer db.Close()

	query := "SELECT id, first_name, last_name, email, username, user_created_at, inactive_status, role FROM execs WHERE 1=1"
	var args []interface{}

	query, args = utils.AddFilters(r, query, args)

	query = utils.AddSorting(r, query)

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("err")
		return nil, utils.ErrorHandler(err, "error retrieving data")
	}
	defer rows.Close()
	for rows.Next() {
		exec := models.Exec{}
		err = rows.Scan(&exec.ID, &exec.FirstName, &exec.LastName, &exec.Email, &exec.Username, &exec.UserCreatedAt, &exec.InactiveStatus, &exec.Role)
		if err != nil {

			return nil, utils.ErrorHandler(err, "error retrieving data")
		}
		execs = append(execs, exec)
	}
	return execs, nil
}

func GetExecByID(id int) (models.Exec, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Exec{}, utils.ErrorHandler(err, "error retrieving data ")
	}
	defer db.Close()

	var exec models.Exec
	err = db.QueryRow("SELECT id, first_name, last_name, email, username, inactive_status, role  FROM execs WHERE id = ?", id).Scan(&exec.ID, &exec.FirstName, &exec.LastName, &exec.Email, &exec.Username, &exec.InactiveStatus, &exec.Role)
	if err == sql.ErrNoRows {
		return models.Exec{}, utils.ErrorHandler(err, "error retrieving data ")
	} else if err != nil {
		fmt.Println(err)
		return models.Exec{}, utils.ErrorHandler(err, "error retrieving data ")
	}
	return exec, nil
}

func AddExecsDBHandler(newExecs []models.Exec) ([]models.Exec, error) {

	fmt.Println("------ AddExecsDBHandler Called -------")

	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "error adding data")
	}
	defer db.Close()
	stmt, err := db.Prepare(utils.GenerateInsertQuery("execs", models.Exec{}))
	if err != nil {
		return nil, utils.ErrorHandler(err, "error adding data")
	}
	defer stmt.Close()
	fmt.Printf("Execs Add Handler")

	addedExecs := make([]models.Exec, len(newExecs))

	for i, newExec := range newExecs {
		values := utils.GetStructValues(newExec)
		fmt.Println(newExec)

		fmt.Println("VALUES:", values)
		res, err := stmt.Exec(values...)

		if err != nil {
			return nil, utils.ErrorHandler(err, "error adding data")
		}
		lastID, err := res.LastInsertId()
		if err != nil {

			return nil, utils.ErrorHandler(err, "error adding data")
		}
		newExec.ID = int(lastID)
		addedExecs[i] = newExec
	}
	return addedExecs, nil
}

func PatchExecs(updates []map[string]interface{}) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		fmt.Println("ID:", idStr)
		if !ok {
			tx.Rollback()

			return utils.ErrorHandler(err, "Invalid id")
		}

		id, err := strconv.Atoi(idStr)
		fmt.Println("ID:", id)
		if err != nil {
			tx.Rollback()
			return utils.ErrorHandler(err, "invalid id")
		}

		var execFromDb models.Exec
		err = db.QueryRow("SELECT id, first_name, last_name, email, username FROM execs WHERE id = ?", id).Scan(&execFromDb.ID, &execFromDb.FirstName, &execFromDb.LastName, &execFromDb.Email, &execFromDb.Username)

		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {

				return utils.ErrorHandler(err, "Exec Not Found")
			}

			return utils.ErrorHandler(err, "error updating data")
		}

		execVal := reflect.ValueOf(&execFromDb).Elem()
		execType := execVal.Type()

		for k, v := range update {
			if k == "id" {
				continue
			}
			for i := 0; i < execVal.NumField(); i++ {
				field := execType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := execVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("cannot convert %v to %v", val.Type(), fieldVal.Type())
							return utils.ErrorHandler(err, "error updating data")
						}
					}
					break
				}
			}
		}
		_, err = tx.Exec("UPDATE execs SET first_name = ?, last_name = ?, email = ?, username = ? WHERE id = ?", execFromDb.FirstName, execFromDb.LastName, execFromDb.Email, execFromDb.Username, execFromDb.ID)
		if err != nil {
			tx.Rollback()

			return utils.ErrorHandler(err, "error updating data")
		}
	}
	err = tx.Commit()
	if err != nil {
		return utils.ErrorHandler(err, "error updating data")
	}
	return nil
}

func PatchOneExec(id int, updates map[string]interface{}) (models.Exec, error) {

	db, err := ConnectDb()
	if err != nil {
		log.Println(err)
		return models.Exec{}, utils.ErrorHandler(err, "error updating data")
	}
	defer db.Close()

	var existingExec models.Exec
	err = db.QueryRow("SELECT id, first_name, last_name, email, username  FROM execs WHERE id = ?", id).Scan(&existingExec.ID, &existingExec.FirstName, &existingExec.LastName, &existingExec.Email, &existingExec.Username)
	if err != nil {
		if err != sql.ErrNoRows {

			return models.Exec{}, utils.ErrorHandler(err, "Exec not Found")
		}
		return models.Exec{}, utils.ErrorHandler(err, "error updating data")
	}
	execVal := reflect.ValueOf(&existingExec).Elem()
	execType := execVal.Type()
	for k, v := range updates {

		for i := 0; i < execVal.NumField(); i++ {

			field := execType.Field(i)

			if field.Tag.Get("json") == k+",omitempty" {
				if execVal.Field(i).CanSet() {
					execVal.Field(i).Set(reflect.ValueOf(v).Convert(execVal.Field(i).Type()))
				}
			}
		}
	}

	_, err = db.Exec("UPDATE execs SET first_name = ?, last_name = ?, email = ?, username = ? WHERE id = ?", existingExec.FirstName, existingExec.LastName, existingExec.Email, existingExec.Username, existingExec.ID)
	if err != nil {
		return models.Exec{}, utils.ErrorHandler(err, "error updating data")
	}
	return existingExec, nil
}

func DeleteOneExec(id int) error {
	db, err := ConnectDb()
	if err != nil {
		log.Println(err)
		return utils.ErrorHandler(err, "error deleting data")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM execs WHERE id = ?", id)
	if err != nil {
		return utils.ErrorHandler(err, "error deleting data")
	}

	fmt.Println(result.RowsAffected())

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "error deleting data")
	}

	if rowsAffected == 0 {
		return utils.ErrorHandler(err, "exec not found")
	}
	return nil
}
```

`internal/models/exec.go`
```go
package models

import "database/sql"

type Exec struct {
	ID                  int            `json:"id,omitempty" db:"id,omitempty"`
	FirstName           string         `json:"first_name,omitempty" db:"first_name,omitempty"`
	LastName            string         `json:"last_name,omitempty" db:"last_name,omitempty"`
	Email               string         `json:"email,omitempty" db:"email,omitempty"`
	Username            string         `json:"username,omitempty" db:"username,omitempty"`
	Password            string         `json:"password,omitempty" db:"password,omitempty"`
	PasswordChangedAt   sql.NullString `json:"password_changed_at,omitempty" db:"password_changed_at,omitempty"`
	UserCreatedAt       sql.NullString `json:"user_created_at,omitempty" db:"user_created_at,omitempty"`
	PasswordResetToken   sql.NullString `json:"password_reset_token,omitempty" db:"password_reset_token,omitempty"`
	PasswordTokenExpires sql.NullString `json:"password_token_expires,omitempty" db:"password_token_expires,omitempty"`
	InactiveStatus      bool           `json:"inactive_status,omitempty" db:"inactive_status,omitempty"`
	Role                string         `json:"role,omitempty" db:"role,omitempty"`
}
```

## Passwords - Hashing

We are not making an API for our software that is publicly available, like instagram or facebook or something like that where anyone can register and login. This is an enterprise software where the user administrators will create a new user. There's no signing up for a new user. The new users are created by the adminitrators. So once an employee joins the executive staff, then that employee will be added to the database and a user will be created for that executive.

When it comes to hashing passwords for secure storage, the choice of hashing algorithms is critical for ensuring both security and efficiency. The three commonly recommended algorithms are Bcrypt, Argon2 and Pbkdf2.

- bcrypt
	- Well established, secure
	- It incorporates a salt to protect against rainbow table attacks and is adaptive, meaning the iteration count can be increased over time to make it slower as computing power increases.
	- Popular

- Argon2
	- Winner of the Password Hashing Competition
	- Three variants: Argon2d, Argon2i and Argon2id
	- Highly efficient
	- Argon2id is recommended for most use-cases as it provides a balance of resistance against both side channel and GPU attacks.
	- In terms of efficiency, Argon2 is highly efficient and allows for fine-tuning of memory usafe, time, cost and parallelism, making it suitable for a wide range of environments.

- PBKDF2 (Password Based Key Derivation Function 2)
	- NIST-approved key derivation function
	- can be slower compared to bcrypt and Argon2 specially when configured with high iteration count for better security.
	- It applies a pseudo random function such as HMAC to the input password along with a salt value and repeats the process many times to produce a derived key. It is considered secure but less resistant to certain types of attacks like side channel attacks.

For our API, we will use Argon2. package: `argon2` (golang.org/x/crypto/argon2)
```bash
go get golang.org/x/crypto/argon2
```


`repository/sqlconnect/execs_crud.go`
```go
func AddExecsDBHandler(newExecs []models.Exec) ([]models.Exec, error) {
#	fmt.Println("------ AddExecsDBHandler Called -------")
#	db, err := ConnectDb()
#	if err != nil {
#		return nil, utils.ErrorHandler(err, "error adding data")
#	}
#	defer db.Close()
#	stmt, err := db.Prepare(utils.GenerateInsertQuery("execs", models.Exec{}))
#	if err != nil {
#		return nil, utils.ErrorHandler(err, "error adding data")
#	}
#	defer stmt.Close()
#	fmt.Printf("Execs Add Handler")
#
#	addedExecs := make([]models.Exec, len(newExecs))
	// previous code 
	for i, newExec := range newExecs {

		if newExec.Password == "" {
			return nil, utils.ErrorHandler(errors.New("password is blank"), "please enter a password")
		}

		salt := make([]byte, 16)
		_, err := rand.Read(salt)
		if err != nil {
			return nil, utils.ErrorHandler(errors.New("failed to generate salt"), "error adding data")
		}

		hash := argon2.IDKey([]byte(newExec.Password), salt, 1, 64*1024, 4, 32)
		saltBase64 := base64.StdEncoding.EncodeToString(salt)
		hashBase64 := base64.StdEncoding.EncodeToString(hash)
		encodedHash := fmt.Sprintf("%s.%s", saltBase64, hashBase64)		
		newExec.Password = encodedHash
		
		// rest of the code in the ADDExecDBHandler

#		values := utils.GetStructValues(newExec)
#		fmt.Println(newExec)
#		fmt.Println("VALUES:", values)
#		res, err := stmt.Exec(values...)
#		if err != nil {
#			return nil, utils.ErrorHandler(err, "error adding data")
#		}
#		lastID, err := res.LastInsertId()
#		if err != nil {
#			return nil, utils.ErrorHandler(err, "error adding data")
#		}
#		newExec.ID = int(lastID)
#		addedExecs[i] = newExec
	}
	return addedExecs, nil
}
```