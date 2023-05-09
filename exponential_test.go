package backoff

import (
	"context"
	. "github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestExponentialBackoff(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	bo := NewExponential().Max(time.Millisecond * 50).Interval(time.Millisecond * 25).Jitter(time.Millisecond * 10).MaxAttempts(5).Init()
	Equal(t, bo.Sleep(ctx, 0), nil)
	Equal(t, bo.Sleep(ctx, 1), nil)
	Equal(t, bo.Sleep(ctx, 2), nil)
	Equal(t, bo.Sleep(ctx, 3), nil)
	Equal(t, bo.Sleep(ctx, 4), nil)
	Equal(t, bo.Sleep(ctx, 5), ErrMaxAttemptsReached)

	cancel()
	Equal(t, bo.Sleep(ctx, 0), context.Canceled)
}
