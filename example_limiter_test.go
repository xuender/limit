package limit_test

import (
	"fmt"
	"time"

	"github.com/xuender/limit"
)

func ExampleLimiter() {
	start := time.Now()
	limiter := limit.NewLimiter(10, time.Second, func(num int) {
		fmt.Println(int(time.Since(start)/(time.Millisecond*10)), num)
	})

	limiter.Add(1)
	limiter.Add(2)
	limiter.Add(3)

	time.Sleep(time.Second)

	// Output:
	// 10 1
	// 20 2
	// 30 3
}
