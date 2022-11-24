package limit_test

import (
	"fmt"
	"time"

	"github.com/xuender/limit"
)

func ExampleLimits() {
	start := time.Now()
	lim := limit.NewLimits(10, time.Second, func(key string, num int) {
		fmt.Println(key, int(time.Since(start)/(time.Millisecond*10)), num)
	})

	lim.SetDefault("A")
	lim.Set("B", 20, time.Second)

	lim.Add("A", 1)
	lim.Add("A", 2)
	lim.Add("B", 3)

	time.Sleep(time.Second)

	// Output:
	// B 5 3
	// A 10 1
	// A 20 2
}
