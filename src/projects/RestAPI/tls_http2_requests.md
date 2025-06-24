# TLS + HTTP2 Requests

- While we test our server of file ([server.go](../../../golang-udemy-code/rest_api/api/server.go)), we see a TLS handshake error while sending our first request. The error in our go https server indicates that there are issues with the SSL and TLS certificate being used. 

The error states TLS Handshake error and then EOF. This error typically indicates that the connection was closed unexpectedly during the TLS handshake.

EOF - End Of File. In the context of network communication and the errors we are encountering, an EOF error typically indicates that one side of the connection has closed the connection unexpectedly or that the connection has been terminated. In general meaning of EOF in networking is connectino termination and protocol errors. So here instead of connection termination, the EOF probably means that there is a protocol error. In the context or TLS or https connection, an EOF error can also occur during the handshake process. If the connection is aborted of if there is a mismatch in expectations between the client and server then we get the EOF error.