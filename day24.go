package adventofcode2022

type Blizzard struct {
	Pos complex128
	Dir complex128
}

type Valley struct {
	Width, Height int
	Start, End    complex128
	Blizzards     []Blizzard
	BlizzardCache map[int]map[complex128]bool
}

func NewValley(lines []string) Valley {
	v := Valley{
		Height:        len(lines),
		Width:         len(lines[0]),
		BlizzardCache: make(map[int]map[complex128]bool),
	}

	for y := range lines {
		for x := range lines[y] {
			pos := R2c(x, y)
			switch lines[y][x] {
			case '.':
				if y == 0 {
					v.Start = pos
				} else if y == len(lines)-1 {
					v.End = pos
				}
			case '^':
				v.Blizzards = append(v.Blizzards, Blizzard{Pos: pos, Dir: -1i})
			case 'v':
				v.Blizzards = append(v.Blizzards, Blizzard{Pos: pos, Dir: 1i})
			case '<':
				v.Blizzards = append(v.Blizzards, Blizzard{Pos: pos, Dir: -1})
			case '>':
				v.Blizzards = append(v.Blizzards, Blizzard{Pos: pos, Dir: 1})
			}
		}
	}

	return v
}

func (v *Valley) blizzardsAt(time int) map[complex128]bool {
	if cached, ok := v.BlizzardCache[time]; ok {
		return cached
	}

	positions := make(map[complex128]bool)
	for _, b := range v.Blizzards {
		// Calculate blizzard position at this time
		pos := b.Pos + complex(float64(time)*real(b.Dir), float64(time)*imag(b.Dir))

		x := int(real(pos))
		y := int(imag(pos))

		// Wrap to inner valley coordinates (1-based to 0-based, apply modulo, back to 1-based)
		innerX := x - 1
		innerY := y - 1
		innerWidth := v.Width - 2
		innerHeight := v.Height - 2

		// Proper modulo that handles negatives correctly
		innerX = ((innerX % innerWidth) + innerWidth) % innerWidth
		innerY = ((innerY % innerHeight) + innerHeight) % innerHeight

		x = innerX + 1
		y = innerY + 1

		positions[R2c(x, y)] = true
	}

	v.BlizzardCache[time] = positions
	return positions
}

func (v *Valley) RenderGrid(time int) []string {
	grid := make([][]byte, v.Height)
	for y := 0; y < v.Height; y++ {
		grid[y] = make([]byte, v.Width)
		for x := 0; x < v.Width; x++ {
			grid[y][x] = '.'
		}
	}

	// Add walls
	for y := 0; y < v.Height; y++ {
		grid[y][0] = '#'
		grid[y][v.Width-1] = '#'
	}
	for x := 0; x < v.Width; x++ {
		grid[0][x] = '#'
		grid[v.Height-1][x] = '#'
	}
	// Start and end openings
	grid[int(imag(v.Start))][int(real(v.Start))] = '.'
	grid[int(imag(v.End))][int(real(v.End))] = '.'

	// Count blizzards at each position
	blizzardCount := make(map[complex128][]complex128)
	for _, b := range v.Blizzards {
		pos := b.Pos + complex(float64(time)*real(b.Dir), float64(time)*imag(b.Dir))
		x := int(real(pos))
		y := int(imag(pos))

		// Wrap to inner valley coordinates (1-based to 0-based, apply modulo, back to 1-based)
		innerX := x - 1
		innerY := y - 1
		innerWidth := v.Width - 2
		innerHeight := v.Height - 2

		// Proper modulo that handles negatives correctly
		innerX = ((innerX % innerWidth) + innerWidth) % innerWidth
		innerY = ((innerY % innerHeight) + innerHeight) % innerHeight

		x = innerX + 1
		y = innerY + 1

		finalPos := R2c(x, y)
		blizzardCount[finalPos] = append(blizzardCount[finalPos], b.Dir)
	}

	// Add blizzards to grid
	for pos, dirs := range blizzardCount {
		x := int(real(pos))
		y := int(imag(pos))
		if len(dirs) > 1 {
			grid[y][x] = byte('0' + len(dirs))
		} else {
			dir := dirs[0]
			if real(dir) == 1 {
				grid[y][x] = '>'
			} else if real(dir) == -1 {
				grid[y][x] = '<'
			} else if imag(dir) == 1 {
				grid[y][x] = 'v'
			} else if imag(dir) == -1 {
				grid[y][x] = '^'
			}
		}
	}

	result := make([]string, v.Height)
	for y := 0; y < v.Height; y++ {
		result[y] = string(grid[y])
	}
	return result
}

func (v *Valley) isValid(pos complex128, time int) bool {
	x := int(real(pos))
	y := int(imag(pos))

	// Check if at start or end position
	if pos == v.Start || pos == v.End {
		return !v.blizzardsAt(time)[pos]
	}

	// Check if within valley bounds (not on walls)
	if x <= 0 || x >= v.Width-1 || y <= 0 || y >= v.Height-1 {
		return false
	}

	// Check if there's a blizzard at this position
	return !v.blizzardsAt(time)[pos]
}

func (v *Valley) shortestPath(start, end complex128, startTime int) int {
	type State struct {
		Pos  complex128
		Time int
	}

	// We begin just before the entrance
	// First move will be into the entrance at time=1
	queue := []State{{Pos: start, Time: startTime}}
	visited := make(map[State]bool)
	visited[State{Pos: start, Time: startTime}] = true

	moves := []complex128{0, 1, -1, 1i, -1i} // wait, right, left, down, up

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Already at end
		if curr.Pos == end {
			return curr.Time
		}

		nextTime := curr.Time + 1

		// Try all possible moves
		for _, move := range moves {
			nextPos := curr.Pos + move

			if v.isValid(nextPos, nextTime) {
				state := State{Pos: nextPos, Time: nextTime}
				if !visited[state] {
					visited[state] = true
					queue = append(queue, state)
				}
			}
		}
	}

	return -1 // No path found
}

func Day24(lines []string, part1 bool) uint {
	v := NewValley(lines)

	if part1 {
		time := v.shortestPath(v.Start, v.End, 0)
		return uint(time)
	}

	// Part 2: go to end, back to start, then to end again
	time1 := v.shortestPath(v.Start, v.End, 0)
	time2 := v.shortestPath(v.End, v.Start, time1)
	time3 := v.shortestPath(v.Start, v.End, time2)
	return uint(time3)
}
