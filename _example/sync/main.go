package main

import (
	"fmt"
	"time"

	"github.com/xuender/limit"
)

func main() {
	count := 0
	limiter := limit.NewSync(1000, time.Second*10)

	go func() {
		for {
			limiter.Wait()
			count++
		}
	}()

	for range time.NewTicker(time.Second).C {
		fmt.Printf("sync QPS: [%d]\t\r", count)
		count = 0
	}
}
