package adventofcode2022

import "fmt"

func Day22(lines []string, part1 bool) int {
	const (
		north = 0 + 1i
		south = 0 - 1i
		east  = 1 + 0i
		west  = -1 + 0i
	)
	var (
		m = make(map[complex128]bool) // 1-based

		walk = func(c complex128) bool {
			// entry and it is true
			b, ok := m[c]
			return ok && b
		}
		wall = func(c complex128) bool {
			// entry and it is false
			b, ok := m[c]
			return ok && !b
		}
		wrap = func(c complex128) bool {
			// no entry
			_, ok := m[c]
			return !ok
		}

		parseMap = func() complex128 {
			var maxX, maxY int
			for y := 0; y < len(lines); y++ {
				if lines[y] == "" {
					break
				}

				// lines are not right padded
				thisX := len(lines[y])
				if thisX > maxX {
					maxX = thisX
				}
				maxY = y

				for x := 0; x < thisX; x++ {
					c := complex(float64(x+1), float64(y+1))
					switch lines[y][x] {
					case '.': // exists and can walk
						m[c] = true
					case '#': // exists and cannot walk
						m[c] = false
					default:
					}
				}
			}
			return complex(float64(maxX+1), float64(maxY+1))
		}

		/*
			set = func(c, facing complex128, n int) {
				for i := 0; i < n; i++ {
					m[c] = false
					c += facing
				}
			}
				border = func(dim complex128) { // create a border around map
					set(0+0i, south, int(imag(dim)))
					set(0+0i, east, int(real(dim)))
					set(dim, north, int(imag(dim)))
					set(dim, west, int(real(dim)))
				}
		*/

		facing = east
	)

	parseMap()
	// border(dim)
	position := func() complex128 {
		c := 1 + 1i
		for {
			if walk(c) {
				break
			}
			if wall(c) {
				break
			}
			c += east
		}
		return c
	}()

	for _, cmd := range NewCommands(lines[len(lines)-1]) {
		// move
		for j := 0; j < cmd.count; {
			p2 := position + facing
			if wrap(p2) {
				// we stepped on a wrap field, carriage return into
				// opposite direction until wrap (or border, which is a
				// wrap)
				opposite := facing * -1
				for walk(p2+opposite) || wall(p2+opposite) {
					p2 += opposite
				}
			}
			if wall(p2) {
				break
			}
			if walk(p2) {
				position = p2
				j++
			}
		}
		facing *= cmd.turn
	}
	result := 1000 * int(imag(position))
	result += 4 * int(real(position))
	switch facing {
	case east:
		// NOP
	case south:
		result += 1
	case west:
		result += 2
	case north:
		result += 3
	}
	return result
}

type Command struct {
	count int
	turn  complex128
}

func NewCommands(s string) (cs []Command) {
	var (
		n      int
		noturn complex128
	)
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n *= 10
			n += int(c - '0')
		} else {
			var turn complex128
			// y extends south, flip
			switch c {
			case 'L':
				turn = 0 - 1i
			case 'R':
				turn = 0 + 1i
			default:
				panic(fmt.Sprintf("Illegal turn %c", c))
			}
			cs = append(cs, Command{n, turn})
			n = 0
		}
	}
	// railing nomber
	cs = append(cs, Command{n, noturn})
	return
}
