//go:build js && wasm

package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	js.Global().Set("formatMakeScript", js.FuncOf(formatMakeScript))
	<-c
}

func formatMakeScript(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return js.ValueOf("Missing input")
	}
	input := args[0].String()
	output := IndentMakeScript(input)
	return js.ValueOf(output)
}


