package utils

import (
	"runtime"
)

const RunningInWASM = runtime.GOARCH == "wasm"

// https://en.wikipedia.org/wiki/Absolute_difference
func AbsDiffInt(x, y int) int {
	return max(x, y) - min(x, y)
}
