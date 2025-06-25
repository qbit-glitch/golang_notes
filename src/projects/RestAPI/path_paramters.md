# Path Params

Path Params are a way to capture values in the URL path to use them within our application. They are commonly used in web applications to pass data between the client and server in clean and readable way.

Path paramters are a fundamental part of building RESTful APIs and web applications in Go.You can handle path parameters using the standard library by manually parsing the URL, and this approach give you complete control over how you extract and use path parameters. However after Go version 1.22, it has become even more easier how we extract path parameters.

```go
