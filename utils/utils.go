package utils

import (
	"runtime"
)

const RunningInWASM = runtime.GOARCH == "wasm"

// https://en.wikipedia.org/wiki/Absolute_difference
func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
