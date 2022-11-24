package limit

import (
	"net/http"
	"time"
)

type data struct {
	writer  http.ResponseWriter
	request *http.Request
}

// LimitHandler is a middleware that performs Limiter given http.Handler struct.
func LimitHandler(qps int, timeOut time.Duration, next http.Handler) http.Handler {
	limiter := NewLimiter(qps, timeOut, func(hand data) { next.ServeHTTP(hand.writer, hand.request) })
	middle := func(w http.ResponseWriter, r *http.Request) {
		if err := limiter.Add(data{writer: w, request: r}); err != nil {
			w.Header().Add("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("request limit."))
		}
	}

	return http.HandlerFunc(middle)
}

// LimitFuncHandler is a middleware that performs Limiter given request handler function.
func LimitFuncHandler(qps int, timeOut time.Duration, nextFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return LimitHandler(qps, timeOut, http.HandlerFunc(nextFunc))
}
