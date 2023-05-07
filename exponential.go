package backoff

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"time"
)

var (
	ErrMaxAttemptsReached = errors.New("max attempts reached")
)

// ExponentialBuilder helps to build the final exponential backoff entity
type ExponentialBuilder struct {
	factor      float64
	interval    time.Duration
	jitter      time.Duration
	max         time.Duration
	maxAttempts int
}

// NewExponential create a new exponential backoff builder with sane defaults.
func NewExponential() ExponentialBuilder {
	return ExponentialBuilder{
		factor:   1.75,
		interval: time.Second,
		jitter:   time.Millisecond * 250,
	}
}

// Factor sets a factor for the backoff algorithm.
func (e ExponentialBuilder) Factor(factor float64) ExponentialBuilder {
	e.factor = factor
	return e
}

// Interval sets base wait interval for the backoff algorithm.
func (e ExponentialBuilder) Interval(interval time.Duration) ExponentialBuilder {
	e.interval = interval
	return e
}

// Jitter sets the maximum jitter for the backoff algorithm.
func (e ExponentialBuilder) Jitter(jitter time.Duration) ExponentialBuilder {
	e.jitter = jitter
	return e
}

// Max sets the maximum timeout despite the number of attempts.
// none/zero is the default.
func (e ExponentialBuilder) Max(max time.Duration) ExponentialBuilder {
	e.max = max
	return e
}

// MaxAttempts sets the maximum number of attempts before the Sleep(...) function begins returning ErrMaxAttemptsReached.
func (e ExponentialBuilder) MaxAttempts(max int) ExponentialBuilder {
	e.maxAttempts = max
	return e
}

// Init returns a read-only(thread safe) Exponential backoff entity for use.
func (e ExponentialBuilder) Init() Exponential {
	return Exponential{
		factor:      e.factor,
		interval:    float64(e.interval),
		jitter:      int64(e.jitter),
		max:         int64(e.max),
		maxAttempts: e.maxAttempts,
	}
}

// Exponential is the final read-only(thread safe) backoff entity
type Exponential struct {
	factor      float64
	interval    float64
	jitter      int64
	max         int64
	maxAttempts int
}

// Duration accepts attempt and returns the backoff duration o sleep for.
func (e Exponential) Duration(attempt int) time.Duration {
	d := math.Pow(e.factor, float64(attempt)) * e.interval
	if e.jitter > 0 {
		d += float64(rand.Int63n(e.jitter))
	}
	i := int64(d)
	if math.IsInf(d, 1) || i < 0 || (e.max > 0 && i > e.max) {
		return time.Duration(e.max)
	}
	return time.Duration(d)
}

// Sleep is a convenience function wrapping Duration and allowing the sleep time to be cancelled via the Context.
//
// This function can also return ErrMaxAttemptsReached if the max attempts have been reached.
func (e Exponential) Sleep(ctx context.Context, attempt int) error {
	if e.maxAttempts > 0 && attempt >= e.maxAttempts {
		return ErrMaxAttemptsReached
	}
	t := time.NewTimer(e.Duration(attempt))
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
