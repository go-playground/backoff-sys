## Package backoff-sys

<img align="right" src="https://raw.githubusercontent.com/go-playground/backoff-sys/master/logo.jpg">![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![Actions Status](https://github.com/go-playground/backoff-sys/workflows/Lint%20%26%20Test/badge.svg)](https://github.com/go-playground/backoff-sys/actions)
[![GoDoc](https://godoc.org/github.com/go-playground/backoff-sys?status.svg)](https://pkg.go.dev/github.com/go-playground/backoff-sys)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package backoff-sys provides the bare building blocks for backing off and can be used to build more complex backoff packages, but this is likely enough.
This includes:
- [x] Exponential backoff, with jitter
- [ ] Linear backoff, with jitter

Example
-------
```go
// go run _examples/exponential/main.go
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/backoff-sys"
)

func main() {
	bo := backoff.NewExponential().Init()
	for i := 0; i < 5; i++ {
		err := fallible()
		if err != nil {
			d := bo.Duration(i)
			fmt.Printf("Waiting: %s\n", d)
			time.Sleep(d)
			continue
		}
	}
}

func fallible() error {
	return errors.New("failed")
}
```

or with cancelable sleep helper

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/backoff-sys"
)

func main() {
	bo := backoff.NewExponential().Init()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 5; i++ {
		err := fallible()
		if err != nil {
			start := time.Now()
			if err := bo.Sleep(ctx, i); err != nil {
				panic(err)
			}
			fmt.Printf("Waited %s\n", time.Since(start))
			continue
		}
	}
}

func fallible() error {
	return errors.New("failed")
}
``` 

License
------
Distributed under MIT License, please see license file in code for more details.
