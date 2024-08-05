package rdb_test

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/limit/rdb"
)

// nolint: testableexamples
func ExampleDistributed() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	start := time.Now()
	limiter := rdb.NewDistributed(client, "key", 1000, time.Second)
	ctx := context.Background()

	_ = limiter.Wait(ctx)
	_ = limiter.Wait(ctx)
	_ = limiter.Wait(ctx)

	fmt.Println(time.Since(start))
}
