package adventofcode2022

import (
	"fmt"
	"io"
)

// WritePBM writes a Portable Bitmap (PBM) in ASCII format (P1).
// The pixel function returns true for black (1) and false for white (0).
func WritePBM(w io.Writer, width, height int, pixel func(x, y int) bool) error {
	// Write PBM header
	_, err := fmt.Fprintf(w, "P1 %d %d\n", width, height)
	if err != nil {
		return err
	}

	// Write pixel data row by row (top to bottom)
	for y := range height {
		for x := range width {
			var bit byte
			if pixel(x, y) {
				bit = '1'
			} else {
				bit = '0'
			}
			_, err := w.Write([]byte{bit})
			if err != nil {
				return err
			}
			if x < width-1 {
				_, err = w.Write([]byte{' '})
				if err != nil {
					return err
				}
			}
		}
		_, err = w.Write([]byte{'\n'})
		if err != nil {
			return err
		}
	}

	return nil
}
