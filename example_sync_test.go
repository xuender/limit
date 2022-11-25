package limit_test

import (
	"fmt"
	"time"

	"github.com/xuender/limit"
)

func ExampleSync() {
	start := time.Now()
	limiter := limit.NewSync(10, time.Second)

	_ = limiter.Wait()
	_ = limiter.Wait()
	_ = limiter.Wait()

	fmt.Println(time.Since(start).Milliseconds() / 10 * 10)

	// Output:
	// 300
}
