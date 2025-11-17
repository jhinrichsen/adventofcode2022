package adventofcode2022

import "image"

type blizzard struct {
	pos image.Point
	dir image.Point
}

type Day24Puzzle struct {
	width, height int
	start, end    image.Point
	blizzards     []blizzard
	blizzardCache map[uint]map[image.Point]bool
}

func NewDay24(lines []string) Day24Puzzle {
	p := Day24Puzzle{
		height:        len(lines),
		width:         len(lines[0]),
		blizzardCache: make(map[uint]map[image.Point]bool),
	}

	for y := range lines {
		for x := range lines[y] {
			pos := image.Pt(x, y)
			switch lines[y][x] {
			case '.':
				if y == 0 {
					p.start = pos
				} else if y == len(lines)-1 {
					p.end = pos
				}
			case '^':
				p.blizzards = append(p.blizzards, blizzard{pos: pos, dir: image.Pt(0, -1)})
			case 'v':
				p.blizzards = append(p.blizzards, blizzard{pos: pos, dir: image.Pt(0, 1)})
			case '<':
				p.blizzards = append(p.blizzards, blizzard{pos: pos, dir: image.Pt(-1, 0)})
			case '>':
				p.blizzards = append(p.blizzards, blizzard{pos: pos, dir: image.Pt(1, 0)})
			}
		}
	}

	return p
}

func (p *Day24Puzzle) blizzardPosition(b blizzard, time uint) image.Point {
	// Calculate position after time steps
	x := b.pos.X + int(time)*b.dir.X
	y := b.pos.Y + int(time)*b.dir.Y

	// Wrap to inner valley coordinates (1-based to 0-based, apply modulo, back to 1-based)
	innerX := x - 1
	innerY := y - 1
	innerWidth := p.width - 2
	innerHeight := p.height - 2

	// Proper modulo that handles negatives correctly
	innerX = ((innerX % innerWidth) + innerWidth) % innerWidth
	innerY = ((innerY % innerHeight) + innerHeight) % innerHeight

	return image.Pt(innerX+1, innerY+1)
}

func (p *Day24Puzzle) blizzardsAt(time uint) map[image.Point]bool {
	if cached, ok := p.blizzardCache[time]; ok {
		return cached
	}

	positions := make(map[image.Point]bool)
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
	grid[p.start.Y][p.start.X] = '.'
	grid[p.end.Y][p.end.X] = '.'

	// Count blizzards at each position using the SAME code as the solver
	blizzardCount := make(map[image.Point][]image.Point)
	for _, b := range p.blizzards {
		pos := p.blizzardPosition(b, time)
		blizzardCount[pos] = append(blizzardCount[pos], b.dir)
	}

	// Add blizzards to grid
	for pos, dirs := range blizzardCount {
		if len(dirs) > 1 {
			grid[pos.Y][pos.X] = byte('0' + len(dirs))
		} else {
			dir := dirs[0]
			if dir.X == 1 {
				grid[pos.Y][pos.X] = '>'
			} else if dir.X == -1 {
				grid[pos.Y][pos.X] = '<'
			} else if dir.Y == 1 {
				grid[pos.Y][pos.X] = 'v'
			} else if dir.Y == -1 {
				grid[pos.Y][pos.X] = '^'
			}
		}
	}

	result := make([]string, p.height)
	for y := 0; y < p.height; y++ {
		result[y] = string(grid[y])
	}
	return result
}

func (p *Day24Puzzle) isValid(pos image.Point, time uint) bool {
	// Check if at start or end position
	if pos == p.start || pos == p.end {
		return !p.blizzardsAt(time)[pos]
	}

	// Check if within valley bounds (not on walls)
	if pos.X <= 0 || pos.X >= p.width-1 || pos.Y <= 0 || pos.Y >= p.height-1 {
		return false
	}

	// Check if there's a blizzard at this position
	return !p.blizzardsAt(time)[pos]
}

func (p *Day24Puzzle) shortestPath(start, end image.Point, startTime uint) uint {
	type state struct {
		pos  image.Point
		time uint
	}

	queue := []state{{pos: start, time: startTime}}
	visited := make(map[state]bool)
	visited[state{pos: start, time: startTime}] = true

	moves := []image.Point{
		{X: 0, Y: 0},  // wait
		{X: 1, Y: 0},  // right
		{X: -1, Y: 0}, // left
		{X: 0, Y: 1},  // down
		{X: 0, Y: -1}, // up
	}

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
			nextPos := curr.pos.Add(move)

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
