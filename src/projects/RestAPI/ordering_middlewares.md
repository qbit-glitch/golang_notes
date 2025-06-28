# Middleware Ordering

Whenever we are making an API, we need to ensure that our middlewares operate in the most logical order and we need to consider how each middleware functions and interacts with the request and response lifecycle. In the API that we made, the order that we chose ensures that each middleware has the opportunity to function correctly, with the innermost middleware handling the request as it passes ourwards through the stack.

```go
secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))

server := &http.Server{
    Addr: port,
    Handler: secureMux,
    TLSConfig: tlsConfig,
}
```