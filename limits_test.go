package limit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
)

func TestNewLimits(t *testing.T) {
	t.Parallel()

	limits := limit.NewLimits(100, time.Second, func(_, _ int) {
		time.Sleep(time.Millisecond)
	})

	assert.NotNil(t, limits.Add(1, 100))
	limits.SetDefault(1, 2)
	limits.SetDefault(2)
	assert.Nil(t, limits.Add(1, 100))
	limits.Set(2, 100, time.Second)
	assert.Nil(t, limits.Add(2, 100))
}
