package limit

import "errors"

var (
	ErrFull    = errors.New("queue full")
	ErrTimeOut = errors.New("time out")
)
