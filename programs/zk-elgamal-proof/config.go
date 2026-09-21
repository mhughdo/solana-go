package zk

import (
	"runtime"
	"sync/atomic"
)

// maxConcurrency is 0 until set, meaning the runtime.NumCPU() default.
var maxConcurrency atomic.Int64

// SetMaxConcurrency caps how many proofs can run concurrently.
// Values below 1 are ignored. It defaults to runtime.NumCPU().
//
// This must be called before any proof is run as a later call has no effect.
func SetMaxConcurrency(n int) {
	if n < 1 {
		return
	}
	maxConcurrency.Store(int64(n))
}

// MaxConcurrency reports the cap in effect, whether set or defaulted.
func MaxConcurrency() int {
	if n := maxConcurrency.Load(); n > 0 {
		return int(n)
	}
	return runtime.NumCPU()
}
