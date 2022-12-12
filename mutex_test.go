package limit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
)

func TestMutex_Wait(t *testing.T) {
	t.Parallel()

	mutex := limit.NewMutex(1, time.Second)
	ass := assert.New(t)

	go func() {
		time.Sleep(time.Millisecond * 300)
		ass.Nil(mutex.Wait())
	}()

	go func() {
		time.Sleep(time.Millisecond * 500)
		ass.NotNil(mutex.Wait())
	}()

	ass.Nil(mutex.Wait())
}

func TestMutex_Try(t *testing.T) {
	t.Parallel()

	mutex := limit.NewMutex(10, time.Second)
	ass := assert.New(t)

	ass.NotNil(mutex.Try())
	time.Sleep(time.Millisecond * 100)
	ass.Nil(mutex.Try())
}

func TestMutex_Timeout(t *testing.T) {
	t.Parallel()

	mutex := limit.NewMutex(1, time.Millisecond)
	ass := assert.New(t)

	ass.NotNil(mutex.Wait())
}
