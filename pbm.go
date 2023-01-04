package adventofcode2022

import (
	"fmt"
	"io"
)

// Portable Bitmap Encoder
// Reference: https://netpbm.sourceforge.net/doc/pbm.html
//
// Sample result for day 10 part 2
//
// ###..###....##..##..####..##...##..###..
// #..#.#..#....#.#..#....#.#..#.#..#.#..#.
// ###..#..#....#.#..#...#..#....#..#.#..#.
// #..#.###.....#.####..#...#.##.####.###..
// #..#.#....#..#.#..#.#....#..#.#..#.#....
// ###..#.....##..#..#.####..###.#..#.#....
//
// or, when using the block character
//
//  ███..███....██..██..████..██...██..███..
//  █..█.█..█....█.█..█....█.█..█.█..█.█..█.
//  ███..█..█....█.█..█...█..█....█..█.█..█.
//  █..█.███.....█.████..█...█.██.████.███..
//  █..█.█....█..█.█..█.█....█..█.█..█.█....
//  ███..█.....██..█..█.████..███.█..█.█....
//
// or, when using PBM
//
// P1 40 6
// 1 1 1 0 0 1 1 1 0 0 0 0 1 1 0 0 1 1 0 0 1 1 1 1 0 0 1 1 0 0 0 1 1 0 0 1 1 1 0 0
// 1 0 0 1 0 1 0 0 1 0 0 0 0 1 0 1 0 0 1 0 0 0 0 1 0 1 0 0 1 0 1 0 0 1 0 1 0 0 1 0
// 1 1 1 0 0 1 0 0 1 0 0 0 0 1 0 1 0 0 1 0 0 0 1 0 0 1 0 0 0 0 1 0 0 1 0 1 0 0 1 0
// 1 0 0 1 0 1 1 1 0 0 0 0 0 1 0 1 1 1 1 0 0 1 0 0 0 1 0 1 1 0 1 1 1 1 0 1 1 1 0 0
// 1 0 0 1 0 1 0 0 0 0 1 0 0 1 0 1 0 0 1 0 1 0 0 0 0 1 0 0 1 0 1 0 0 1 0 1 0 0 0 0
// 1 1 1 0 0 1 0 0 0 0 0 1 1 0 0 1 0 0 1 0 1 1 1 1 0 0 1 1 1 0 1 0 0 1 0 1 0 0 0 0
//

// BitProvider is a callback function called by WritePBM.
// x ranges from 0 to width -1, y from 0 to height -1.
// Both X and Y axis are mathematical, especially Y is _not_ upside down.
// The calling order is guaranteed y height -> 0, and x 0 -> width.
//
// Y ^
//   |
//   |
//   |
//   |
// 0 +------------>
//   0            X
//
type BitProvider func(x, y int) bool

func WritePBM(out io.Writer, width, height int, bits BitProvider) error {
	var buf [2]byte
	_, err := fmt.Fprintf(out, "P1 %d %d\n", width, height)
	if err != nil {
		return err
	}
	// for now, just print one by one, no StringBuilder, no buffer.
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if bits(x, y) {
				buf[0] = '1'
			} else {
				buf[0] = '0'
			}
			if x < width-1 {
				buf[1] = ' '
			} else {
				buf[1] = '\n'
			}
			_, err := out.Write(buf[:])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
