package limit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
)

func TestNewLimiter(t *testing.T) {
	t.Parallel()

	lim := limit.NewLimiter(1, time.Second, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	assert.NotNil(t, lim.Add(1))

	lim = limit.NewLimiter(1, time.Millisecond, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	assert.NotNil(t, lim.Add(1))

	lim = limit.NewLimiter(1000, time.Millisecond, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	time.Sleep(time.Millisecond * 2)
	assert.Nil(t, lim.Add(1))
	lim.Close()
}

func TestLimiterAdd(t *testing.T) {
	t.Parallel()

	lim := limit.NewLimiter(0, time.Second, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
}
