package adventofcode2022

type blizzard struct {
	pos complex128
	dir complex128
}

type Day24Puzzle struct {
	width, height int
	start, end    complex128
	blizzards     []blizzard
	blizzardCache map[uint]map[complex128]bool
}

func NewDay24(lines []string) Day24Puzzle {
	p := Day24Puzzle{
		height:        len(lines),
		width:         len(lines[0]),
		blizzardCache: make(map[uint]map[complex128]bool),
	}

	for y := range lines {
		for x := range lines[y] {
			pos := R2c(x, y)
			switch lines[y][x] {
			case '.':
				if y == 0 {
					p.start = pos
				} else if y == len(lines)-1 {
					p.end = pos
				}
			case '^':
				p.blizzards = append(p.blizzards, blizzard{pos: pos, dir: -1i})
			case 'v':
				p.blizzards = append(p.blizzards, blizzard{pos: pos, dir: 1i})
			case '<':
				p.blizzards = append(p.blizzards, blizzard{pos: pos, dir: -1})
			case '>':
				p.blizzards = append(p.blizzards, blizzard{pos: pos, dir: 1})
			}
		}
	}

	return p
}

func (p *Day24Puzzle) blizzardPosition(b blizzard, time uint) complex128 {
	pos := b.pos + complex(float64(time)*real(b.dir), float64(time)*imag(b.dir))
	x := int(real(pos))
	y := int(imag(pos))

	// Wrap to inner valley coordinates (1-based to 0-based, apply modulo, back to 1-based)
	innerX := x - 1
	innerY := y - 1
	innerWidth := p.width - 2
	innerHeight := p.height - 2

	// Proper modulo that handles negatives correctly
	innerX = ((innerX % innerWidth) + innerWidth) % innerWidth
	innerY = ((innerY % innerHeight) + innerHeight) % innerHeight

	x = innerX + 1
	y = innerY + 1

	return R2c(x, y)
}

func (p *Day24Puzzle) blizzardsAt(time uint) map[complex128]bool {
	if cached, ok := p.blizzardCache[time]; ok {
		return cached
	}

	positions := make(map[complex128]bool)
	for _, b := range p.blizzards {
		positions[p.blizzardPosition(b, time)] = true
	}

	p.blizzardCache[time] = positions
	return positions
}

func (p *Day24Puzzle) renderGrid(time uint) []string {
	grid := make([][]byte, p.height)
	for y := 0; y < p.height; y++ {
		grid[y] = make([]byte, p.width)
		for x := 0; x < p.width; x++ {
			grid[y][x] = '.'
		}
	}

	// Add walls
	for y := 0; y < p.height; y++ {
		grid[y][0] = '#'
		grid[y][p.width-1] = '#'
	}
	for x := 0; x < p.width; x++ {
		grid[0][x] = '#'
		grid[p.height-1][x] = '#'
	}
	// Start and end openings
	grid[int(imag(p.start))][int(real(p.start))] = '.'
	grid[int(imag(p.end))][int(real(p.end))] = '.'

	// Count blizzards at each position using the SAME code as the solver
	blizzardCount := make(map[complex128][]complex128)
	for _, b := range p.blizzards {
		pos := p.blizzardPosition(b, time)
		blizzardCount[pos] = append(blizzardCount[pos], b.dir)
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

	result := make([]string, p.height)
	for y := 0; y < p.height; y++ {
		result[y] = string(grid[y])
	}
	return result
}

func (p *Day24Puzzle) isValid(pos complex128, time uint) bool {
	x := int(real(pos))
	y := int(imag(pos))

	// Check if at start or end position
	if pos == p.start || pos == p.end {
		return !p.blizzardsAt(time)[pos]
	}

	// Check if within valley bounds (not on walls)
	if x <= 0 || x >= p.width-1 || y <= 0 || y >= p.height-1 {
		return false
	}

	// Check if there's a blizzard at this position
	return !p.blizzardsAt(time)[pos]
}

func (p *Day24Puzzle) shortestPath(start, end complex128, startTime uint) uint {
	type state struct {
		pos  complex128
		time uint
	}

	queue := []state{{pos: start, time: startTime}}
	visited := make(map[state]bool)
	visited[state{pos: start, time: startTime}] = true

	moves := []complex128{0, 1, -1, 1i, -1i} // wait, right, left, down, up

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Already at end
		if curr.pos == end {
			return curr.time
		}

		nextTime := curr.time + 1

		// Try all possible moves
		for _, move := range moves {
			nextPos := curr.pos + move

			if p.isValid(nextPos, nextTime) {
				s := state{pos: nextPos, time: nextTime}
				if !visited[s] {
					visited[s] = true
					queue = append(queue, s)
				}
			}
		}
	}

	return 0 // No path found
}

func Day24(puzzle Day24Puzzle, part1 bool) uint {
	if part1 {
		return puzzle.shortestPath(puzzle.start, puzzle.end, 0)
	}

	// Part 2: go to end, back to start, then to end again
	time1 := puzzle.shortestPath(puzzle.start, puzzle.end, 0)
	time2 := puzzle.shortestPath(puzzle.end, puzzle.start, time1)
	time3 := puzzle.shortestPath(puzzle.start, puzzle.end, time2)
	return time3
}
