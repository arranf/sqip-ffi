package main

import "C"
import (
	"github.com/denisbrodbeck/sqip"
)

//export MakeSVG
func MakeSVG(path string) *C.char {
	workSize := 256  // large images get resized to this - higher size grants no boons
	count := 8       // number of primitive SVG shapes
	mode := 1        // shape type
	alpha := 128     // alpha value
	repeat := 0      // add N extra shapes each iteration with reduced search (mostly good for beziers)
	workers := 1     // number of parallel workers
	background := "" // background color (hex)

	svg, _, _, err := sqip.Run(path, workSize, count, mode, alpha, repeat, workers, background)
	if err != nil {
		result := "Error " + err.Error()
		return C.CString(result)
	}
	return C.CString(svg)
}

func main() {}
