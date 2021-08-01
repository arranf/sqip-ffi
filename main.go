package main

import "C"
import (
	"github.com/denisbrodbeck/sqip"
)

//export MakeSVG
func MakeSVG(path string, numberOfPrimitives int, mode int, alpha int, workers int) *C.char {
	workSize := 256  // large images get resized to this - higher size grants no boons
	repeat := 0      // add N extra shapes each iteration with reduced search (mostly good for beziers)
	background := "" // background color (hex)

	svg, _, _, err := sqip.Run(path, workSize, numberOfPrimitives, mode, alpha, repeat, workers, background)
	if err != nil {
		result := "Error " + err.Error()
		return C.CString(result)
	}
	return C.CString(svg)
}

func main() {}
