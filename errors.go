package limit

import "errors"

var (
	ErrTimeOut = errors.New("time out")
	ErrKey     = errors.New("key not exist")
	ErrQPS     = errors.New("qps <= 0")
)
