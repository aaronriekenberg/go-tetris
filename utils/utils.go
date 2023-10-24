package utils

import (
	"runtime"

	"golang.org/x/exp/constraints"
)

func RunningInWASM() bool {
	return runtime.GOARCH == "wasm"
}

func IntegerAbs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
