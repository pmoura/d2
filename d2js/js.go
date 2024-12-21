//go:build js && wasm

package main

import (
	"syscall/js"

	"oss.terrastruct.com/d2/d2js/d2wasm"
)

func main() {
	api := d2wasm.NewD2API()

	api.Register("getParentID", d2wasm.GetParentID)
	api.Register("getObjOrder", d2wasm.GetObjOrder)
	api.Register("getRefRanges", d2wasm.GetRefRanges)
	api.Register("compile", d2wasm.Compile)
	api.Register("getBoardAtPosition", d2wasm.GetBoardAtPosition)
	api.Register("encode", d2wasm.Encode)
	api.Register("decode", d2wasm.Decode)
	api.Register("version", d2wasm.GetVersion)

	api.ExportTo(js.Global())

	if cb := js.Global().Get("onWasmInitialized"); !cb.IsUndefined() {
		cb.Invoke()
	}
	select {}
}
