package limit

import "time"

// Sync channel based sync rate limit.
type Sync struct {
	limiter *Async[chan<- struct{}]
}

// NewSync returns a new sync rate limit.
func NewSync(qps int, timeOut time.Duration) *Sync {
	limiter := NewAsync(qps, timeOut, func(end chan<- struct{}) {
		end <- struct{}{}
	})

	return &Sync{limiter: limiter}
}

// Wait and returns a error.
func (p *Sync) Wait() error {
	end := make(chan struct{})

	if err := p.limiter.Add(end); err != nil {
		return err
	}

	<-end

	return nil
}

// Close channel.
func (p *Sync) Close() {
	p.limiter.Close()
}
