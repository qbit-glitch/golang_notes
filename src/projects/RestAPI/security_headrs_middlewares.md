# Security Headers Middleware

Applying Security Headers makes a significant difference in securing our API and improving browser behavior. They mitigate risks associated with common web vulnerabilities and ensure that your application adheres to best practices for web-security. Without these headers, our API will be more vulnerables to attacks and could lead to compromised security for your users. So applying these security headers in your REST API can significantly enhance the security of your application by mitigating various attack vectors. These are actual security headers which enhance the security of your API and they do by protecting your API against various attack vectors.


```go
package middlewares

import "net/http"

func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode-block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("X-Powered-By","Django")

		w.Header().Set("Server", "")
		w.Header().Set("X-Permitterd-Cross-Domain-Policies", "none")
		w.Header().Set("Cache-Control", "no-store, no-chache, must-revalidate, max-age=0")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy","require-corp")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Permissions-Policy", "geolocation=(self), microphone=()")

		next.ServeHTTP(w,r)
	})
}
```