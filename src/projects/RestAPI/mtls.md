# mTLS

Use a common openssl.conf file to generate the key.pem and cert.pem files. Use this command:
```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -config openssl.cnf
```

`openssl.conf`
```cnf
[req]
default_bits        = 2048
distinguished_name  = req_distinguished_name
req_extensions      = req_ext
prompt              = no

[req_distinguished_name]
C = US
ST = State
L = City
O = Organization
OU = Organization Unit
CN = localhost

[req_ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
DNS.2 = 127.0.0.1
```

## mTLS

mTLS stands for mutual TLS is an extension of TLS that requires both the client and the server to authenticate each other using certificates. This is more commonly used in environments where both parties need to establish a higher level of trust, such as in internal communications between microservices or in specific client server applications.

mTLS is typically not used for public facing websites like bank portals as it requires the client, in this case our web-browser or device to have a client certificate installed which is not practical for general consumer use. It's used in desktop or mobile applications for validating the certificate between the client and the server. So this is an enhanced level of security that we can implement.

Add this code snippet in the `tlsConfig` variable in `server.go`:
```go
// Configure the TLS
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS12,
    ClientAuth: tls.RequireAndVerifyClientCert,
    ClientCAs: loadClientCAs(),
}
```

Make another function `loadClientCAs()`
```go

func loadClientCAs() *x509.CertPool{
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Could not load the client CA:", err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}
```

[Final Code](../../../golang-udemy-code/rest_api/mtls_server/mtls_server.go)