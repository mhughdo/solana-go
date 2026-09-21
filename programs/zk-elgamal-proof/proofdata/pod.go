package proofdata

import (
	"bytes"
	"fmt"
	"unsafe"
)

// podBytes copies v's memory into a fresh buffer.
func podBytes[T any](v *T) []byte {
	return bytes.Clone(unsafe.Slice((*byte)(unsafe.Pointer(v)), unsafe.Sizeof(*v)))
}

// podRead copies exactly sizeof(T) bytes into v's memory.
func podRead[T any](v *T, b []byte) error {
	if uintptr(len(b)) != unsafe.Sizeof(*v) {
		return fmt.Errorf("zk: got %d bytes, want %d", len(b), unsafe.Sizeof(*v))
	}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(v)), len(b)), b)
	return nil
}
