package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/xuender/limit"
)

func main() {
	http.Handle("/", limit.FuncHandler(1, time.Second*3, ping))
	// nolint
	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "PONG")
	fmt.Println(time.Now().Format("15:04:05.00"))
}
