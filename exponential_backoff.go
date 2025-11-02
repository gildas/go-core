package core

import (
	"math"
	"math/rand/v2"
	"time"
)

// ExponentialBackoff returns a duration to wait before retrying an operation, using an exponential backoff algorithm.
//
// The delay is calculated as initialDelay * 2^(attempt-1), with a maximum of maxDelay. Attempt is 1-based (i.e. the first attempt is 1).
//
// A jitter factor is applied to the delay to avoid thundering herd problems. The jitter is a random value between
// -jitter * delay and +jitter * delay.
func ExponentialBackoff(attempt int, initialDelay time.Duration, maxDelay time.Duration, jitter float64) time.Duration {
	delay := min(time.Duration(math.Pow(2, float64(attempt-1)))*initialDelay, maxDelay)
	return delay + time.Duration(rand.Float64()*float64(delay)*jitter*2) - time.Duration(float64(delay)*jitter)
}
