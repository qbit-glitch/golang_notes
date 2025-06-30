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
