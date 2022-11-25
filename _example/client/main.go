package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	group := sync.WaitGroup{}
	group.Add(100)

	for i := 0; i < 100; i++ {
		go func(num int) {
			res, _ := http.Get("http://127.0.0.1:8080")

			fmt.Printf("%02d => %d, %v\n", num, res.StatusCode, time.Since(start))
			group.Done()
		}(i)
	}

	group.Wait()
}
