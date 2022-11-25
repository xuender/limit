package limit_test

import (
	"fmt"
	"time"

	"github.com/xuender/limit"
)

func ExampleAsync() {
	start := time.Now()
	limiter := limit.NewAsync(10, time.Second, func(num int) {
		fmt.Println(time.Since(start).Milliseconds()/10*10, num)
	})

	_ = limiter.Add(1)
	_ = limiter.Add(2)
	_ = limiter.Add(3)

	time.Sleep(time.Second)

	// Output:
	// 100 1
	// 200 2
	// 300 3
}
