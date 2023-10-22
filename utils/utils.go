package utils

import "runtime"

func RunningInWASM() bool {
	return runtime.GOARCH == "wasm"
}
