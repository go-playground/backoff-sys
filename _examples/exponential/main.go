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
