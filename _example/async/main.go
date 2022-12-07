package main

import (
	"fmt"
	"time"

	"github.com/xuender/limit"
)

func main() {
	count := 0
	limiter := limit.NewAsync(1000, time.Second*10, func(i int) {
		count++
	})

	go func() {
		for {
			limiter.Add(1)
		}
	}()

	for range time.NewTicker(time.Second).C {
		fmt.Printf("async QPS: [%d]\t\r", count)
		count = 0
	}
}
