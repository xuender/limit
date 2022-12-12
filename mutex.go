package limit

import (
	"sync"
	"time"
)

// Mutex rate limit.
type Mutex struct {
	interval time.Duration
	last     time.Time
	mutex    sync.Mutex
	timeout  time.Duration
}

// NewMutex returns a new mutex rate limit.
//
// qps less than 1 unlimited.
//
// Play: https://go.dev/play/p/ogcvT7o4ENI
func NewMutex(qps int, timeOut time.Duration) *Mutex {
	return &Mutex{
		interval: time.Second / time.Duration(qps),
		mutex:    sync.Mutex{},
		last:     time.Now(),
		timeout:  timeOut,
	}
}

// Wait and returns a error.
func (p *Mutex) Wait() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if wait := p.waiting(); wait > 0 {
		if wait > p.timeout {
			return ErrTimeOut
		}

		time.Sleep(wait)

		p.last = p.last.Add(p.interval)
	}

	return nil
}

func (p *Mutex) waiting() time.Duration {
	dru := time.Since(p.last)

	if sleep := p.interval - dru; sleep > 0 {
		return sleep
	}

	p.last = p.last.Add(dru)

	return 0
}

// Try and returns a error.
func (p *Mutex) Try() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.waiting() > 0 {
		return ErrWait
	}

	return nil
}
