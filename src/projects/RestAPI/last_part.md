# Part - {Last}

## Why `WHERE 1=1` in the SQL query

`WHERE 1=1` is always true and acts as a placeholder to simplify appending additional conditions dynamically. It allows us to add conditions without worrying about whether it's the first conditions or not, avoiding the need to check if you should start with WHERE or add an AND. 

`WHERE 1=1` also simplify the code for adding filters. Without  `WHERE 1=1`, you would need to check if the WHERE clause already exists before adding `AND` for each filter which adds extra complexity. And using this `WHERE 1=1`, we can actually handle multiple filters.

In conclusion, using `WHERE 1=1` is a common practice for making dynamic query buliding more straightformward. It avoids the need for complex conditional logic when appending multiple filter conditions, making your code cleaner and easier to maintain.


## Advanced Filtering Technique: GET - Getting entries based on multiple criteria

```go
func addFilters(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for param, dbField := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			query += " AND " + dbField + " = ?"
			args = append(args, value)
		}
	}
	return query, args
}
```


## Advanced Sort Order Technique: GET - Sorting nad getting entries based on multiple criteria

`.Get()` returns a single value. This method returns the first value associated with the `sortby` key in the query parameters as a string. If there are multiple `sortby` parameters, only the first one is returned and if the key is not present it returns an empty string.

On the other hand, `r.URL.Query()["sortby"]` returns a slice of values. This method returns all values associated with the `sortby` key as a slice of strings. If there are multiple `sortby` parameters, all of them are returned. And all of them are returned in a slice form. If the key is not present, it returns a nil slice.

So in our case, where we want to handle multiple sorting criteria, `r.URL.Query()["sortby"]` is more appropriate because it gives us all the sorting parameters as a slice, allowing us to iterate through them and apply multiple sorting conditions.

```go
func addSorting(r *http.Request, query string) string {
	sortParams := r.URL.Query()["sortby"]
	if len(sortParams) > 0 {
		query += " ORDER BY"
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				continue
			}
			field, order := parts[0], parts[1]

			if !isValidField(field) || !isValidOrder(order) {
				continue
			}
			if i > 0 {
				query += ","
			}
			query += " " + field + " " + order
		}
	}
	return query
}

func isValidOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func isValidField(field string) bool {
	validFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"class":      true,
		"email":      true,
		"subject":    true,
	}
	return validFields[field]
}
```


## Updating a complete entry - PUT

PUT and PATCH differ from one another in the way that PUT is meant to completely replace the entity. However PATCH is just modifying the entry slightly. If there is a minor modification then in that case we use patch. However, PUT is not used that much.

When we are posting data, we use `Exec()`. When we are retrieving data, we use `Query()`. In `PUT` request, we need to send all the values for all the fields. If we send a blank value, then the black value will be updated.

```go

func updateTeacherHandler(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var updatedTeacher models.Teacher
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return 
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT * FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.Class, &existingTeacher.Email, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Subject)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to Retrieve Data", http.StatusInternalServerError)
		return
	}

	updatedTeacher.ID = existingTeacher.ID
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		http.Error(w,"Error Updating teacher", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTeacher)

}
```

## Modifying an Entry - PATCH

```go
// PATCH /teachers/{id}
func patchTeachersHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, class, email, first_name, last_name, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.Class, &existingTeacher.Email, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Subject)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Unable to Retrieve Data", http.StatusInternalServerError)
		return
	}

	// Apply updates
	for k, v := range updates {
		switch k {
		case "first_name":
			existingTeacher.FirstName = v.(string)
		case "last_name":
			existingTeacher.LastName = v.(string)
		case "email":
			existingTeacher.Email = v.(string)
		case "class":
			existingTeacher.Class = v.(string)
		case "subject":
			existingTeacher.Subject = v.(string)
		}
	}

	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", existingTeacher.FirstName, existingTeacher.LastName, existingTeacher.Email, existingTeacher.Class, existingTeacher.Subject, existingTeacher.ID)
	if err != nil {
		http.Error(w, "Error Updating teacher", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingTeacher)

}
```

## Improving our PATCH functionality - `reflect` package

Instead of using switch statements for every field, use the `reflect` package. 

```go
// // Apply updates
// for k, v := range updates {
// 	switch k {
// 	case "first_name":
// 		existingTeacher.FirstName = v.(string)
// 	case "last_name":
// 		existingTeacher.LastName = v.(string)
// 	case "email":
// 		existingTeacher.Email = v.(string)
// 	case "class":
// 		existingTeacher.Class = v.(string)
// 	case "subject":
// 		existingTeacher.Subject = v.(string)
// 	}
// }

// Apply updates using `reflect` package
teacherVal := reflect.ValueOf(&existingTeacher).Elem()
teacherType := teacherVal.Type()
for k, v := range updates {
    for i := 0; i < teacherVal.NumField(); i++ {
        
        field := teacherType.Field(i)

        if field.Tag.Get("json") == k + ",omitempty" {
            if teacherVal.Field(i).CanSet() {
             teacherVal.Field(i).Set(reflect.ValueOf(v).Convert(teacherVal.Field(i).Type()))
            }
        }
    }
}
```


## Deleting an entry - DELETE

```go

func deleteTeachersHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return
	}

	fmt.Println(result.RowsAffected())

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error retrieving delete result", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	// Response Body -> Optional
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Teacher deleted successfully",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)

	// Return status of NoContent -> Compulsory
	w.WriteHeader(http.StatusNoContent)

}
```


## Modernizing Routes:

With Go version 1.22, we can extract path parameters just like we used to extract query parameters. We can just mention the method right before the route and pass it the handler function associates with that HTTP method.

Path with `...` with the ellipsis, will include all the values after the intial route as the path parameter associated with the key that we have mentioned before the three dots `...`.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Method based routing
	mux.HandleFunc("POST /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Item created")
	})

	// Method based routing
	mux.HandleFunc("DELETE /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Item deleted")
	})

	// Wildcard in pattern - path parameter
	mux.HandleFunc("GET /teachers/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Teacher ID: %s", r.PathValue("id"))
	})

	// Wildcard with "...."
	mux.HandleFunc("/files/{path...}", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Path: %s", r.PathValue("path"))
	})

	// Confusion when there are two conflicting wildcards
	mux.HandleFunc("/path1/{param1}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Param1: %s", r.PathValue("param1"))
	})
	// mux.HandleFunc("/{param2}/path2", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Param2: %s", r.PathValue("param2"))
	// })
	/* Error:
		/{param2}/path2 and /files/{path...} both match some paths, like "/files/path2".
        But neither is more specific than the other.
        /{param2}/path2 matches "/param2/path2", but /files/{path...} doesn't.
        /files/{path...} matches "/files/", but /{param2}/path2 doesn't.
	*/

	mux.HandleFunc("/path2/param2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Nothing to see here")
	})

	http.ListenAndServe(":8080", mux)
}
```

## Refactoring MUX

There should me only a single space between the http method and your route. If there is no space or more than one space, our application will not recognize the `DELETE` method associates with this route.

```go
mux.HandleFunc("DELETE /teachers/", handlers.TeachersHandler)
```
For now we are focusing on only the `/teachers/` route. So the `router.go` will now look like this:
```go
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
```
