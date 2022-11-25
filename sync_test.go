package limit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/limit"
)

func TestSync_Wait(t *testing.T) {
	t.Parallel()

	sync := limit.NewSync(1, time.Second)
	ass := assert.New(t)

	go func() {
		time.Sleep(time.Millisecond * 300)
		ass.Nil(sync.Wait())
	}()

	go func() {
		time.Sleep(time.Millisecond * 500)
		ass.NotNil(sync.Wait())
	}()

	ass.Nil(sync.Wait())
}

func TestSync_Close(t *testing.T) {
	t.Parallel()

	sync := limit.NewSync(1, time.Second)

	sync.Close()
}
