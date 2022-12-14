package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/limit"
)

func main() {
	count := 0
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	limiter := limit.NewRdb(client, "limit_test", 1000, time.Second)

	go func() {
		for {
			limiter.Wait()
			count++
		}
	}()

	for range time.NewTicker(time.Second).C {
		fmt.Printf("redis QPS: [%d]   \r", count)
		count = 0
	}
}
