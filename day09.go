package adventofcode2022

import (
	"math"
	"math/cmplx"
	"strconv"
	"strings"
)

func Day09(lines []string) int {
	var directions = map[byte]complex128{
		'U': 0 + 1i,
		'R': 1 + 0i,
		'D': 0 - 1i,
		'L': -1 + 0i,
	}

	reach := cmplx.Abs(1 + 1i)
	var h, t complex128

	// next step
	step := func() complex128 {
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

	ts := make(map[complex128]bool)

	// record current tail position
	ts[t] = true

	for _, line := range lines {
		fields := strings.Fields(line)
		c := directions[fields[0][0]]
		steps, _ := strconv.Atoi(fields[1])

		for steps > 0 {
			// head step
			h += c

			// tail step if too far away
			d := cmplx.Abs(h - t)
			if d > reach {
				inc := step()
				t += inc
				ts[t] = true
			}

			steps--
		}
	}
	return len(ts)
}
