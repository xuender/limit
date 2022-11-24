package limit

import "time"

const _timeOut = time.Microsecond * 100

type Limit[T any] struct {
	interval time.Duration
	last     time.Time
	elems    chan T
	yield    func(T)
}

func NewLimit[T any](qps int, timeOut time.Duration, yield func(T)) *Limit[T] {
	res := &Limit[T]{yield: yield}

	if qps > 0 {
		size := qps * int(timeOut/time.Second)

		if size < 1 {
			size = 1
		}

		res.elems = make(chan T, size)
		res.interval = time.Second / time.Duration(qps)
		res.last = time.Now()

		go res.limit()
	}

	return res
}

func (p *Limit[T]) limit() {
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

func (p *Limit[T]) Close() {
	close(p.elems)
}

func (p *Limit[T]) Add(elem T) error {
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
