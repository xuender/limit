package limit

import "time"

const _timeOut = time.Microsecond * 100

// Async channel based async rate limit.
type Async[T any] struct {
	interval time.Duration
	last     time.Time
	elems    chan T
	yield    func(T)
}

// NewAsync returns a new async rate limit.
func NewAsync[T any](qps int, timeOut time.Duration, yield func(T)) *Async[T] {
	limiter := &Async[T]{yield: yield}

	if qps > 0 {
		size := qps * int(timeOut/time.Second)

		if size < 1 {
			size = 1
		}

		limiter.elems = make(chan T, size)
		limiter.interval = time.Second / time.Duration(qps)
		limiter.last = time.Now()

		go limiter.run()
	}

	return limiter
}

func (p *Async[T]) run() {
	for elem := range p.elems {
		dur := time.Since(p.last)

		if sleep := p.interval - dur; sleep > 0 {
			dur = p.interval

			time.Sleep(sleep)
		}

		p.last = p.last.Add(dur)
		p.yield(elem)
	}
}

// Add elem returns a error.
func (p *Async[T]) Add(elem T) error {
	if p.interval == 0 {
		p.yield(elem)

		return nil
	}

	select {
	case p.elems <- elem:
		return nil
	case <-time.After(_timeOut):
		return ErrTimeOut
	}
}

// Close channel.
func (p *Async[T]) Close() {
	close(p.elems)
}
