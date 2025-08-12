package adventofcode2022

func Day23(lines []string, rounds int) int {
	var (
		m        = make(map[complex128]bool)
		parseMap = func() {
			for y := len(lines) - 1; y >= 0; y-- {
				for x := len(lines[y]) - 1; x >= 0; x-- {
					if lines[y][x] == '#' {
						// invert y
						m[R2c(x, len(lines)-1-y)] = true
					}
				}
			}
		}
		any = func() complex128 {
			for k := range m {
				return k
			}
			panic("want at least one map entry but map is empty")
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
		countEmpty = func(min, max complex128) int {
			return int((1+real(max)-real(min))*(1+imag(max)-imag(min))) - len(m)
		}

		/*
			dump = func(round int) {
				min, max := rect()
				if round == 0 {
					fmt.Printf("== Initial State (%v) ==\n", max-min)
				} else {
					fmt.Printf("== End of Round %d (%vx%v -> %v) ==\n", round, min, max, max-min)
				}
				for y := imag(max); y >= imag(min); y-- {
					for x := real(min); x <= real(max); x++ {
						if m[complex(x, y)] {
							fmt.Printf("%c", '#')
						} else {
							if x == 0 && y == 0 {
								fmt.Printf("%c", '+')
							} else {
								fmt.Printf("%c", '.')
							}
						}
					}
					fmt.Println()
				}
				fmt.Println()
			}
		*/

		proposals = []struct {
			preds []complex128
			move  complex128
		}{
			{[]complex128{North, NorthEast, NorthWest}, North},
			{[]complex128{South, SouthEast, SouthWest}, South},
			{[]complex128{East, NorthEast, SouthEast}, East},
			{[]complex128{West, NorthWest, SouthWest}, West},
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

	// dump(0)
	for round := 1; round <= 10; round++ {

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
		if stable {
			break
		}

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
		shiftProposals()
		// dump(round)
	}

	return countEmpty(rect())
}
