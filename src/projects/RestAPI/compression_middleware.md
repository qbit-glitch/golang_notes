# Compression Middleware

Using compression middleware in Go can be very beneficial for imptroving the performance of your web applications. Compression reduces the size of the response body sent over the network, which can significantly decrease loading times for your application. This is especially important for large assets like images, stylesheets and javascript files. By compressing responses you can minimize the amount of data transferred over the network, reducing bandwidth costs and improving overall efficiency.

Why use compression middleware ?
- Improved Performance
- Reduced Bandwidth usage
- Better User Experience
- Easy Integration

If our application primarily serves small payloads, for example a small API response, the overhead of compression may not be worth the gain as the compression ratio for small data may be minimal. Another factor to consider is that compression requires CPU resources. If your server is already under heavy load, adding compression may lead to performance degradation.

When you might not need Compression Middleware:
- Small Payloads
- Already Compressed Assets
- CPU Overhead

In conclusion, implementing compression middleware in your Go application can lead to significant performance improvements by reducing response sizes, lowering bandwidth usage, and enhancing the user experience. It can be easily integrated into your http server setup and is especially useful for applications serving large or static content.

```go
package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip"){
			next.ServeHTTP(w,r)
		}

		// Set the response header
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Wrap the ResponseWriter
		w = &gzipResponseWriter{
			ResponseWriter: w,
			Writer: gz,
		}

		next.ServeHTTP(w,r)
		fmt.Println("Sent from Compression Middleware")

	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error){
	return g.Writer.Write(b)
}
```

