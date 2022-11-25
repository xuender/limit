package limit

import (
	"net/http"
	"time"
)

// Handler is a middleware that performs Limiter given http.Handler struct.
func Handler(qps int, timeOut time.Duration, next http.Handler) http.Handler {
	limiter := NewSync(qps, timeOut)
	middle := func(writer http.ResponseWriter, request *http.Request) {
		if err := limiter.Wait(); err != nil {
			writer.Header().Add("Content-Type", "text/plain; charset=utf-8")
			writer.WriteHeader(http.StatusTooManyRequests)
			_, _ = writer.Write([]byte("request limit."))

			return
		}

		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(middle)
}

// FuncHandler is a middleware that performs Limiter given request handler function.
func FuncHandler(qps int, timeOut time.Duration, nextFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return Handler(qps, timeOut, http.HandlerFunc(nextFunc))
}
