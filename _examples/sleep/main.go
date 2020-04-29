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
