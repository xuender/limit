package limit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
)

func TestNewLimiter(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	limiter := limit.NewAsync(1, time.Second, func(num int) { time.Sleep(time.Microsecond) })

	ass.Nil(limiter.Add(1))
	ass.Nil(limiter.Add(1))
	ass.Equal(1, limiter.Len())
	ass.NotNil(limiter.Add(1))
	limiter.Close()

	limiter = limit.NewAsync(1, time.Millisecond, func(num int) { time.Sleep(time.Microsecond) })

	ass.Nil(limiter.Add(1))
	ass.Nil(limiter.Add(1))
	ass.NotNil(limiter.Add(1))
	limiter.Close()

	limiter = limit.NewAsync(1000, time.Millisecond, func(num int) { time.Sleep(time.Microsecond) })

	ass.Nil(limiter.Add(1))
	ass.Nil(limiter.Add(1))
	time.Sleep(time.Millisecond * 2)
	ass.Nil(limiter.Add(1))
	limiter.Close()
}

func TestLimiterAdd(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	limiter := limit.NewAsync(0, time.Second, func(num int) { time.Sleep(time.Microsecond) })

	defer limiter.Close()

	ass.Nil(limiter.Add(1))
	ass.Nil(limiter.Add(1))
	ass.Nil(limiter.Add(1))
}
