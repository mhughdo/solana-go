package zk

import (
	"runtime"
	"testing"
)

func TestMaxConcurrency(t *testing.T) {
	t.Cleanup(func() { maxConcurrency.Store(0) })

	if got := MaxConcurrency(); got != runtime.NumCPU() {
		t.Errorf("default = %d, want NumCPU (%d)", got, runtime.NumCPU())
	}

	SetMaxConcurrency(0)
	if got := MaxConcurrency(); got != runtime.NumCPU() {
		t.Errorf("after SetMaxConcurrency(0) = %d, want NumCPU (%d)", got, runtime.NumCPU())
	}

	SetMaxConcurrency(2)
	if got := MaxConcurrency(); got != 2 {
		t.Errorf("after SetMaxConcurrency(2) = %d, want 2", got)
	}
}
