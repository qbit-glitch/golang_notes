# HTTPs certificates - SSL/TLS

The `.pem` extension stands for Privacy Enhanced Mail. It is a base64 encoded DER certificates. So PEM files are commonly used for storing and transmitting cryptographic keys, certificates and other data. 

DER certificate means Distinguished Encoding Rules and it is a binary encoding for `X509` certificates.

The PEM formate is base64 encoding making it easier to read and transport in text based protocols. So even if we are using http1.1 when we are using text-based protocols, it is easier to transport thse certificate files. pem files have specific headers and footers to identify the type of content they hold.

The server use the `key.pem` file to prove its identity and establish a secure connection. And then `cert.pem` file is provided to clients to verify server's identity and to encrypt the data sent to the server. 

In summary, the PEM files are just text files containing cryptographic keys and certificates encoded in base64 with specific headers and footers and `keep.pem` contains the private key for decryption and signing, while `cert.pem` contains the public key and certificate for encryption and identity verification. These files are essential for setting up https, where the server must prove it's identity to clients and establish a secure communication channel.

### Generate key and certificate separately

1. To generate the key separately use this command :
```bash
openssl genpkey -algorithm RSA -out server.key -pkeyopt rsa_keygen_bits:2048
```

2. To generate a certificate based on the above key:
```bash
openssl req -new -x509 -key server.key -out server.crt -days365
```

