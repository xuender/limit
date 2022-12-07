package limit_test

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/limit"
)

// nolint: testableexamples
func ExampleRdb() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	start := time.Now()
	limiter := limit.NewRdb(client, "key", 1000)

	_ = limiter.Wait()
	_ = limiter.Wait()
	_ = limiter.Wait()

	fmt.Println(time.Since(start))
}
