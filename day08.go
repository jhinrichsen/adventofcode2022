package adventofcode2022

func Day08(lines []string) int {
	// for sparse fields i like to use reals as index, this field is full
	// no clue for any advanced stuff
	// in case of doubt choose brute force
	// -- me

	X := len(lines[0])
	Y := len(lines)
	visible := func(x, y, dx, dy int) bool {
		h := lines[y][x]
		for x > 0 && x < X-1 && y > 0 && y < Y-1 {
			x += dx
			y += dy
			if lines[y][x] >= h {
				return false
			}
		}
		return true
	}
	var count int
	for y := 1; y < Y-1; y++ {
		for x := 1; x < X-1; x++ {
			// north
			if visible(x, y, 0, -1) {
				count++
				continue
			}

			// east
			if visible(x, y, 1, 0) {
				count++
				continue
			}

			// south
			if visible(x, y, 0, 1) {
				count++
				continue
			}

			// west
			if visible(x, y, -1, 0) {
				count++
				continue
			}
		}
	}
	// outside are always visible
	outside := 2*X + 2*Y - 4 // 4 corners
	return count + outside
}
