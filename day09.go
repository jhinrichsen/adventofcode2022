package adventofcode2022

import (
	"math"
	"math/cmplx"
	"strconv"
	"strings"
)

func Day09(lines []string, knots []complex128) int {
	var directions = map[byte]complex128{
		'U': 0 + 1i,
		'R': 1 + 0i,
		'D': 0 - 1i,
		'L': -1 + 0i,
	}

	reach := cmplx.Abs(1 + 1i)

	ts := make(map[complex128]bool) // tail positions

	// record current tail position
	ts[knots[len(knots)-1]] = true

	for _, line := range lines {
		fields := strings.Fields(line)
		c := directions[fields[0][0]]

		for steps, _ := strconv.Atoi(fields[1]); steps > 0; steps-- {
			// head step
			knots[0] += c

			for i := 1; i < len(knots); i++ {
				// tail step if too far away
				d := cmplx.Abs(knots[i] - knots[i-1])
				if d > reach {
					inc := step(knots[i-1], knots[i])
					knots[i] += inc
				}
			}
			ts[knots[len(knots)-1]] = true
		}
	}
	return len(ts)
}

// step tail towards head
func step(h, t complex128) complex128 {
	dx := real(h) - real(t)
	dy := imag(h) - imag(t)

	// limit to one step
	if dx != 0 {
		dx /= math.Abs(dx)
	}
	if dy != 0 {
		dy /= math.Abs(dy)
	}
	return complex(dx, dy)
}
