# Add HTTP2 and HTTPs to our API

In order to incorporate https functionality on our server we need to use certificate. These certificates are usually issued by certifying authorities like Google Trust Services or Cloudflare. We don't need certificates as of now coz we are in development phase and we are not putting our API in production. During deploying our API into production, we would need certificates from a certifying authority because those are widely accepted by the browsers. 

## Steps : 

1. As of now we will just generate self-signed certificates on your computer. Command :

```bash
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
```

2. We configure out server using those certificates that we generated. And we define the port and the handlers of the orders route and the users route. 

```go
// Load the TLS cert and key
cert := "cert.pem"
key := "key.pem"

// Configure the TLS
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS12,
}

// Create a custom server
server := &http.Server{
    Addr: fmt.Sprintf(":%d", port),
    Handler: nil,
    TLSConfig: tlsConfig,
}
```