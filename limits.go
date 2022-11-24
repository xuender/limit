package limit

import (
	"sync"
	"time"
)

type Limits[K comparable, V any] struct {
	limits  map[K]*Limiter[V]
	qps     int
	timeOut time.Duration
	yield   func(K, V)
	mutex   sync.RWMutex
}

func NewLimits[K comparable, V any](qps int, timeOut time.Duration, yield func(K, V)) *Limits[K, V] {
	return &Limits[K, V]{
		limits:  map[K]*Limiter[V]{},
		qps:     qps,
		timeOut: timeOut,
		yield:   yield,
	}
}

func (p *Limits[K, V]) SetDefault(keys ...K) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, key := range keys {
		if limit, has := p.limits[key]; has {
			limit.Close()
		}

		p.limits[key] = NewLimiter(p.qps, p.timeOut, func(elem V) { p.yield(key, elem) })
	}
}

func (p *Limits[K, V]) Set(key K, qps int, timeOut time.Duration) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if limit, has := p.limits[key]; has {
		limit.Close()
	}

	p.limits[key] = NewLimiter(qps, timeOut, func(elem V) { p.yield(key, elem) })
}

func (p *Limits[K, V]) Add(key K, elem V) error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if limit, has := p.limits[key]; has {
		return limit.Add(elem)
	}

	return ErrKey
}
