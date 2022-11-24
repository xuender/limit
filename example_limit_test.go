package limit_test

import (
	"fmt"
	"time"

	"github.com/xuender/limit"
)

func ExampleLimit() {
	now := time.Now()
	lim := limit.NewLimit(100, time.Second, func(num int) {
		fmt.Println(int(time.Since(now)/time.Millisecond), num)
	})

	lim.Add(1)
	lim.Add(2)
	lim.Add(3)

	time.Sleep(time.Millisecond * 100)

	// Output:
	// 10 1
	// 20 2
	// 30 3
}
