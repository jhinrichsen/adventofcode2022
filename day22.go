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

		wall = func(c complex128) bool {
			_, ok := m[c]
			return !ok
		}
		walk = func(c complex128) bool {
			return m[c]
		}
		wrap = func(c complex128) bool {
			return !(wall(c) || walk(c))
		}

		parseMap = func() {
			for y := 0; y < len(lines); y++ {
				if lines[y] == "" {
					break
				}
				for x := 0; x < len(lines[y]); x++ {
					c := complex(float64(x+1), float64(y+1))
					switch lines[y][x] {
					case '.': // exists and can walk
						m[c] = true
					case ' ': // exists and wraps
						m[c] = false
					case '#': // cannot walk
					default:
						panic("illegal tile")
					}
				}
			}
		}

		facing = east

		commands = lines[len(lines)-1]
	)

	parseMap()
	position := func() complex128 {
		c := 1 + 1i
		for {
			if walk(c) {
				break
			}
			c += east
		}
		return c
	}()

	fmt.Printf("starting position: %v\n", position)
	fmt.Printf("wall %t, walk %t, wrap %t\n", wall(position),
		walk(position), wrap(position))
	fmt.Printf("facing %v\n", facing)
	fmt.Printf("commands: %v\n", commands)

	return 0
}
