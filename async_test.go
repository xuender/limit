package limit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
)

func TestNewLimiter(t *testing.T) {
	t.Parallel()

	lim := limit.NewAsync(1, time.Second, func(num int) { time.Sleep(time.Microsecond) })
	ass := assert.New(t)

	ass.Nil(lim.Add(1))
	ass.Nil(lim.Add(1))
	ass.Equal(1, lim.Len())
	ass.NotNil(lim.Add(1))

	lim = limit.NewAsync(1, time.Millisecond, func(num int) { time.Sleep(time.Microsecond) })

	ass.Nil(lim.Add(1))
	ass.Nil(lim.Add(1))
	ass.NotNil(lim.Add(1))

	lim = limit.NewAsync(1000, time.Millisecond, func(num int) { time.Sleep(time.Microsecond) })

	ass.Nil(lim.Add(1))
	ass.Nil(lim.Add(1))
	time.Sleep(time.Millisecond * 2)
	ass.Nil(lim.Add(1))
	lim.Close()
}

func TestLimiterAdd(t *testing.T) {
	t.Parallel()

	lim := limit.NewAsync(0, time.Second, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
}
