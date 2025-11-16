package adventofcode2022

import (
	"strings"
)

func Day10(lines []string, part1 bool) (uint, []string) {
	var (
		// like strconv.Atoi() but... custom. Why not.
		parse = func(s string) int {
			n := 0
			sign := 1
			i := 5
			if s[i] == '-' {
				sign = -1
				i++
			}
			for ; i < len(s); i++ {
				n = 10*n + int(s[i]-'0')
			}
			return sign * n
		}
		dim    = 2 * len(lines) // 2 states/ op max
		states = make([]int, dim)

		cycle = 1 // 1-based
		x     = 1
	)

	for _, line := range lines {
		// cycle 1

		states[cycle] = x
		cycle++
		if line == "noop" {
			continue
		}

		// cycle 2

		states[cycle] = x
		// value changes _after_ cycle
		cycle++
		x += parse(line)
	}

	if part1 {
		sum := 0
		for cycle := 20; cycle <= 220; cycle += 40 {
			sum += cycle * states[cycle]
		}
		return uint(sum), []string{}
	}

	var crt strings.Builder
	const (
		width  = 40
		height = 6
	)
	for cycle := 1; cycle <= width*height; cycle++ {
		position := (cycle - 1) % width
		lit := states[cycle]-1 == position ||
			states[cycle] == position ||
			states[cycle]+1 == position
		if lit {
			crt.WriteByte('#')
		} else {
			crt.WriteByte('.')
		}
		if cycle%width == 0 {
			crt.WriteByte(' ')
		}
	}
	return 0, strings.Fields(crt.String())
}
