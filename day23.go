package adventofcode2022

func Day23(lines []string, part1 bool) uint {
	var (
		m        = make(map[complex128]bool)
		parseMap = func() {
			for y := range lines {
				for x := range lines[y] {
					if lines[y][x] == '#' {
						// Use inverted Y since North is +1i (up)
						m[R2c(x, len(lines)-1-y)] = true
					}
				}
			}
		}
		any = func() complex128 {
			for k := range m {
				return k
			}
			return 0 // Empty map, return origin
		}
		rect = func() (complex128, complex128) {
			// initialize on arbitrary position
			c := any()
			minX, maxX := real(c), real(c)
			minY, maxY := imag(c), imag(c)
			for c = range m {
				if real(c) < minX {
					minX = real(c)
				} else if real(c) > maxX {
					maxX = real(c)
				}
				if imag(c) < minY {
					minY = imag(c)
				} else if imag(c) > maxY {
					maxY = imag(c)
				}
			}
			return complex(minX, minY), complex(maxX, maxY)
		}
		neighbors = func(c complex128) int {
			count := 0
			for _, c2 := range []complex128{
				North, NorthEast, East, SouthEast,
				South, SouthWest, West, NorthWest,
			} {
				if m[c+c2] {
					count++
				}
			}
			return count
		}
		countEmpty = func(min, max complex128) uint {
			return uint((1+real(max)-real(min))*(1+imag(max)-imag(min))) - uint(len(m))
		}

		proposals = []struct {
			preds []complex128
			move  complex128
		}{
			{[]complex128{North, NorthEast, NorthWest}, North},
			{[]complex128{South, SouthEast, SouthWest}, South},
			{[]complex128{West, NorthWest, SouthWest}, West},
			{[]complex128{East, NorthEast, SouthEast}, East},
		}
		shiftProposals = func() {
			prop0 := proposals[0]
			for i := 1; i < len(proposals); i++ {
				proposals[i-1] = proposals[i]
			}
			proposals[len(proposals)-1] = prop0
		}
	)

	parseMap()

	maxRounds := 10
	if !part1 {
		maxRounds = 1000000 // Large number for "run until stable"
	}

	for round := 1; round <= maxRounds; round++ {
		// first half
		dsts := make(map[complex128]complex128, len(m))
		counts := make(map[complex128]int, len(m))
		for c := range m {
			elveStable := neighbors(c) == 0
			if elveStable {
				continue
			}
			for _, proposal := range proposals {
				p0 := proposal.preds[0]
				p1 := proposal.preds[1]
				p2 := proposal.preds[2]
				if !m[c+p0] &&
					!m[c+p1] &&
					!m[c+p2] {

					c2 := c + proposal.move
					dsts[c] = c2
					counts[c2]++
					break
				}
			}
		}

		// second half
		stable := len(dsts) == 0
		if stable && !part1 {
			return uint(round)
		}

		// Only execute moves if not stable
		if !stable {
			for from := range m {
				into, wantsToMove := dsts[from]
				if wantsToMove {
					canMove := counts[into] == 1
					if canMove {
						delete(m, from)
						m[into] = true
					}
				}
			}
		}
		shiftProposals()
	}

	return countEmpty(rect())
}
