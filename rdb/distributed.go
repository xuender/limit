package rdb

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xuender/limit"
)

// Distributed is redis limiter.
type Distributed struct {
	client   redis.Cmdable
	key      string
	old      int64
	interval time.Duration
	last     time.Time
	mutex    sync.Mutex
	timeOut  time.Duration
}

// NewDistributed returns redis limiter.
func NewDistributed(client redis.Cmdable, key string, qps int, timeOut time.Duration) *Distributed {
	return &Distributed{
		client:   client,
		key:      key,
		old:      0,
		interval: time.Second / time.Duration(qps),
		mutex:    sync.Mutex{},
		timeOut:  timeOut,
	}
}

// Wait is concurrent flow control.
func (p *Distributed) Wait(ctx context.Context) error {
	if p.interval <= 0 {
		return limit.ErrQPS
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	num, err := p.client.Incr(ctx, p.key).Result()
	if err != nil {
		return err
	}

	if p.old <= 0 {
		p.old = num
		p.last = time.Now()

		return nil
	}

	interval := p.interval * time.Duration(num-p.old)
	dur := time.Since(p.last)
	sleep := interval - dur

	if sleep >= p.timeOut {
		if _, err := p.client.Decr(context.Background(), p.key).Result(); err != nil {
			return err
		}

		return limit.ErrTimeOut
	}

	p.old = num

	switch {
	case sleep > 0:
		time.Sleep(sleep)
	case sleep < 0 && sleep*-1 > time.Millisecond*3:
		p.last = time.Now()

		return nil
	}

	p.last = p.last.Add(dur + sleep)

	return nil
}
