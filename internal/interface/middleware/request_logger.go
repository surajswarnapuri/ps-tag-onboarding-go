package middleware

import (
	"log"
	"net/http"
	"time"
)

// RequestLogger is a middleware that logs the HTTP request path and basic metadata.
func RequestLogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		// Propagate the context with span to the next handler
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf(
			`{"level":"info","msg":"request completed","method":"%s","path":"%s","duration_ms":%d}`,
			r.Method, r.URL.Path, duration.Milliseconds(),
		)
	})
}
