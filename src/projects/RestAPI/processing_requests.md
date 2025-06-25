# Processing Requests


## Code for parsing the form data

**Parse form data when the form is x-wwww-form-urlencoded**
```go
// Parse form data (necessary for x-www-form-urlencoded)
err := r.ParseForm()
if err != nil {
    http.Error(w, "Error parsing form:", http.StatusBadRequest)
    return
}
fmt.Println("Form:", r.Form)

// Parse Response Data
response := make(map[string]interface{})
for key, value := range r.Form {
    response[key] = value[0]
}
fmt.Println("Processed Response Map:", response)
```



**Parse form data when the form is passed as RAW JSON**
```go
// RAW Body
body, err := io.ReadAll(r.Body)
if err != nil {
    http.Error(w, "Error parsing form:", http.StatusBadRequest)
    return
}
defer r.Body.Close()

fmt.Println("Raw Body:", string(body))


// If you expect JSON data, then Unmarshall it using struct
var userInstance1 User
err = json.Unmarshal(body, &userInstance1)
if err != nil {
    http.Error(w, "Error parsing form:", http.StatusBadRequest)
    return
}
fmt.Println(userInstance1)

// Using maps to Unmarshal the data
userInstance2 := make(map[string]interface{})
err = json.Unmarshal(body, &userInstance2)
if err != nil {
    http.Error(w, "Error parsing form:", http.StatusBadRequest)
    return
}
fmt.Println("Unmarshaled JSON into a map", userInstance2)
```

**Different Options that we can use with the request**

```go
fmt.Println("Body:", r.Body)
fmt.Println("Form:", r.Form)
fmt.Println("Header:", r.Header)
fmt.Println("Context:", r.Context())
fmt.Println("Content Length:", r.ContentLength)
fmt.Println("Host:", r.Host)
fmt.Println("Method:", r.Method)
fmt.Println("Protocol:", r.Proto)
fmt.Println("Remote Addr:", r.RemoteAddr)
fmt.Println("Request URI:", r.RequestURI)
fmt.Println("TLS:", r.TLS)
fmt.Println("Trailer:", r.Trailer)
fmt.Println("Transfer Encoding:", r.TransferEncoding)
fmt.Println("URL:", r.URL)
fmt.Println("User Agent", r.UserAgent())
fmt.Println("Port:", r.URL.Port())
```

