# Limit

[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/limit)](https://goreportcard.com/report/github.com/xuender/limit)
[![tag](https://img.shields.io/github/tag/xuender/limit.svg)](https://github.com/xuender/limit/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://godoc.org/github.com/xuender/limit?status.svg)](https://pkg.go.dev/github.com/xuender/limit)
![Build Status](https://github.com/xuender/limit/actions/workflows/go.yml/badge.svg)
[![Coverage](https://img.shields.io/codecov/c/github/xuender/limit)](https://codecov.io/gh/xuender/limit)
[![License](https://img.shields.io/github/license/xuender/limit)](./LICENSE)

Golang channel based rate limiter.

* supports asynchronous and synchronous calls.
* simple middleware to rate limit HTTP requests.
* requests that may timeout will returns an error immediately.
* call order.
* redis based distributed limiter.

## 💡 Usage

You can import limit using:

```go
import "github.com/xuender/limit"
```

### Async

```go
start := time.Now()
limiter := limit.NewAsync(10, time.Second, func(num int) {
  fmt.Println(time.Since(start), num)
})

_ = limiter.Add(1)
_ = limiter.Add(2)
_ = limiter.Add(3)

time.Sleep(time.Second)

// Output:
// 100ms 1
// 200ms 2
// 300ms 3
```

[[play](https://go.dev/play/p/W9gYA_109Vz)]

### Sync

```go
start := time.Now()
limiter := limit.NewSync(10, time.Second)

_ = limiter.Wait()
_ = limiter.Wait()
_ = limiter.Wait()

fmt.Println(time.Since(start))

// Output:
// 300ms
```

[[play](https://go.dev/play/p/tFrkT_j1obb)]

### Rdb

redis based distributed limiter.

```go
client := redis.NewClient(&redis.Options{
  Addr:     "localhost:6379",
  Password: "",
  DB:       0,
})
start := time.Now()
limiter := limit.NewRdb(client, "key", 1000, time.Second)

_ = limiter.Wait()
_ = limiter.Wait()
_ = limiter.Wait()

fmt.Println(time.Since(start))
```

[[play]](https://go.dev/play/p/SgAqdQRYhiK)

### Handler

```go
http.Handle("/", limit.FuncHandler(1000, time.Second*10,
  func(w http.ResponseWriter, r *http.Request) {
    _, _ = io.WriteString(w, "PONG")
  },
))
http.ListenAndServe(":8080", nil)
```

[[play](https://go.dev/play/p/oAoIZynIdkn)]

## 📝 License

© ender, 2022~time.Now

[MIT License](https://github.com/xuender/limit/blob/master/LICENSE)
