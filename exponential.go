package backoff

import (
	"math"
	"math/rand"
	"time"
)

// ExponentialBuilder helps to build the final exponential backoff entity
type ExponentialBuilder struct {
	factor   float64
	interval time.Duration
	jitter   time.Duration
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

// Init returns a read-only(thread safe) Exponential backoff entity for use.
func (e ExponentialBuilder) Init() Exponential {
	return Exponential{
		factor:   e.factor,
		interval: float64(e.interval),
		jitter:   int64(e.jitter),
	}
}

// Exponential is the final read-only(thread safe) backoff entity
type Exponential struct {
	factor   float64
	interval float64
	jitter   int64
}

// Duration accepts attempt and returns the backoff duration o sleep for.
func (e Exponential) Duration(attempt int) time.Duration {
	d := int64(math.Pow(e.factor, float64(attempt)) * e.interval)
	if e.jitter <= 0 {
		return time.Duration(d)
	}
	jitter := rand.Int63n(e.jitter)
	return time.Duration(d + jitter)
}
