package main

import (
	"net/http"
	"sync"
)

func main() {
	group := sync.WaitGroup{}
	group.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			http.Get("http://127.0.0.1:8080")
			group.Done()
		}()
	}

	group.Wait()
}
