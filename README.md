# limit

[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/limit)](https://goreportcard.com/report/github.com/xuender/limit)

Golang channel based rate limiter.

## use

```go
start := time.Now()
limiter := limit.NewLimiter(10, time.Second, func(num int) {
  fmt.Println(time.Since(start), num)
})

limiter.Add(1)
limiter.Add(2)
limiter.Add(3)

time.Sleep(time.Second)
```

## License

© ender, 2022~time.Now

[MIT License](https://github.com/xuender/limit/blob/master/License)
