# Using Curl to make http2 request

Let's add some some more functionality to our server to log the type of http protocol, if it is actually the http version if it's 1.1 or if it's http2. And also we are going to log the TLS version if it's TLS 1.2 or TLS 1.3

```go

func logRequestDetails(r *http.Request){
    httpVersion := r.Proto
	fmt.Println("Received request with HTTP Version:", httpVersion)
	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS version:", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS version"
	}
}
```

### Curl Commands:

```bash
curl -v "https://localhost:3000/orders"
```
This will result in an error coz it's a self-signed certificate. 

```bash
curl -v -k "https://localhost:3000/orders"
```
`-k` flag is used to ignore the self-signed certificate error. It will bypass the certificate validation.


Sample output

```bash
* Host localhost:3000 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:3000...
* Connected to localhost (::1) port 3000
* ALPN: curl offers h2,http/1.1
* (304) (OUT), TLS handshake, Client hello (1):
* (304) (IN), TLS handshake, Server hello (2):
* (304) (IN), TLS handshake, Unknown (8):
* (304) (IN), TLS handshake, Certificate (11):
* (304) (IN), TLS handshake, CERT verify (15):
* (304) (IN), TLS handshake, Finished (20):
* (304) (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / AEAD-CHACHA20-POLY1305-SHA256 / [blank] / UNDEF
* ALPN: server accepted h2
* Server certificate:
*  subject: C=AU; ST=Non Existent; L=Randome; O=API inc; OU=API Inc; CN=API Inc; emailAddress=test@test.com
*  start date: Jun 24 05:57:36 2025 GMT
*  expire date: Jun 24 05:57:36 2026 GMT
*  issuer: C=AU; ST=Non Existent; L=Randome; O=API inc; OU=API Inc; CN=API Inc; emailAddress=test@test.com
*  SSL certificate verify result: self signed certificate (18), continuing anyway.
* using HTTP/2
* [HTTP/2] [1] OPENED stream for https://localhost:3000/orders
* [HTTP/2] [1] [:method: GET]
* [HTTP/2] [1] [:scheme: https]
* [HTTP/2] [1] [:authority: localhost:3000]
* [HTTP/2] [1] [:path: /orders]
* [HTTP/2] [1] [user-agent: curl/8.7.1]
* [HTTP/2] [1] [accept: */*]
> GET /orders HTTP/2
> Host: localhost:3000
> User-Agent: curl/8.7.1
> Accept: */*
> 
* Request completely sent off
< HTTP/2 200 
< content-type: text/plain; charset=utf-8
< content-length: 24
< date: Tue, 24 Jun 2025 13:12:52 GMT
< 
* Connection #0 to host localhost left intact
Handling incoming orders⏎
```

- ALPN - Application Layer Protocol Negotiation. It's a protocol negotiation to negotiate the protocol that the client and the server will be using. And in the above output.. they have agreed to use http2 protocol.