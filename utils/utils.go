package utils

import (
	"runtime"

	"golang.org/x/exp/constraints"
)

func RunningInWASM() bool {
	return runtime.GOARCH == "wasm"
}

func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
