# Middlewares

Middleware is like a checkpoint or a gatekeeper that stands between a client request and the final processing of that request by the server. It can inspect modify or log the request before it reaches the final destination and it can do the same with the response before it is send back to the client.

A middleware performs a task on the request and then allows the request to move forward onto the next middleware and then from one middleware to the next and then to the next until it finally reaches the handler function.

Middleware in an API serves various purposes:
- Logging
- Authentication and Authorization
- Data Validation
- Error Handling

Mechanism of a middleware :
In Go, middleware is a function that wraps around another function, the actual request handler and this wrapper function can fo something before and or after calling the actual handler.

Concept of `next` in middlewares :
- `http.Handler` interface
```go
type Handler interface{
    ServeHTTP(ResponseWriter, *Request)
}
```
- Middleware Pattern
- Chaining Handlers


### Structure of a Middleware :
```go
func MiddlwareName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w,r)
	})
}
```