package limit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
)

func TestNewLimit(t *testing.T) {
	t.Parallel()

	lim := limit.NewLimit(1, time.Second, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	assert.NotNil(t, lim.Add(1))

	lim = limit.NewLimit(1, time.Millisecond, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	assert.NotNil(t, lim.Add(1))

	lim = limit.NewLimit(1000, time.Millisecond, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	time.Sleep(time.Millisecond * 2)
	assert.Nil(t, lim.Add(1))
	lim.Close()
}

func TestLimitAdd(t *testing.T) {
	t.Parallel()

	lim := limit.NewLimit(0, time.Second, func(num int) { time.Sleep(time.Microsecond) })

	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
	assert.Nil(t, lim.Add(1))
}
