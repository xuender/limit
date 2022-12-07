package limit

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// Rdb is redis limiter.
type Rdb struct {
	client   redis.Cmdable
	key      string
	old      int64
	interval time.Duration
	last     time.Time
	mutex    sync.Mutex
}

// NewRdb returns redis limiter.
func NewRdb(client redis.Cmdable, key string, qps int) *Rdb {
	return &Rdb{
		client:   client,
		key:      key,
		old:      0,
		interval: time.Second / time.Duration(qps),
		mutex:    sync.Mutex{},
	}
}

// Wait is concurrent flow control.
func (p *Rdb) Wait() error {
	if p.interval <= 0 {
		return ErrQPS
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	num, err := p.client.Incr(context.Background(), p.key).Result()
	if err != nil {
		return err
	}

	if p.old <= 0 {
		p.old = num
		p.last = time.Now()

		return nil
	}

	interval := p.interval * time.Duration(num-p.old)
	p.old = num
	dur := time.Since(p.last)
	sleep := interval - dur

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
