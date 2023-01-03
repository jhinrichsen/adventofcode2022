package adventofcode2022

func Day08(lines []string, part1 bool) int {
	if part1 {
		return day08Part1(lines)
	}
	return day08Part2(lines)
}

func day08Part1(lines []string) int {
	// for sparse fields i like to use reals as index, this field is full
	// no clue for any advanced stuff
	//
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

func day08Part2(lines []string) int {
	var (
		X              = len(lines[0])
		Y              = len(lines)
		maxScenicScore = 0
		height         = func(x, y int) int {
			return int(lines[y][x] - '0')
		}
	)
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			var cn int
			h0 := height(x, y)
			// north
			for y2 := y - 1; ; y2-- {
				if y2 < 0 {
					break
				}
				cn++
				h := height(x, y2)
				if h >= h0 {
					break
				}
			}
			// east
			var ce int
			for x2 := x + 1; ; x2++ {
				if x2 >= X {
					break
				}
				ce++
				h := height(x2, y)
				if h >= h0 {
					break
				}
			}
			// south
			var cs int
			for y2 := y + 1; ; y2++ {
				if y2 >= Y {
					break
				}
				cs++
				h := height(x, y2)
				if h >= h0 {
					break
				}
			}
			// west
			var cw int
			for x2 := x - 1; ; x2-- {
				if x2 < 0 {
					break
				}
				cw++
				h := height(x2, y)
				if h >= h0 {
					break
				}
			}

			scenicScore := cn * ce * cs * cw
			if maxScenicScore < scenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}
