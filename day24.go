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

		// Wrap around walls (walls are at y=0, y=Height-1, x=0, x=Width-1)
		x := int(real(pos))
		y := int(imag(pos))

		// Wrap X coordinate (left/right wrapping)
		if x < 1 {
			x = ((x-1)%(v.Width-2) + (v.Width - 2)) + 1
		} else if x >= v.Width-1 {
			x = ((x-1)%(v.Width-2))+1
		}

		// Wrap Y coordinate (up/down wrapping)
		if y < 1 {
			y = ((y-1)%(v.Height-2) + (v.Height - 2)) + 1
		} else if y >= v.Height-1 {
			y = ((y-1)%(v.Height-2))+1
		}

		positions[R2c(x, y)] = true
	}

	v.BlizzardCache[time] = positions
	return positions
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

	// Start just outside the entrance
	initialPos := start - 1i // One position above/before the start
	queue := []State{{Pos: initialPos, Time: startTime}}
	visited := make(map[State]bool)
	visited[State{Pos: initialPos, Time: startTime}] = true

	moves := []complex128{0, 1, -1, 1i, -1i} // wait, right, left, down, up

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Check if we reached the exit (one step past the end)
		if curr.Pos == end+1i {
			return curr.Time
		}

		nextTime := curr.Time + 1

		// Try all possible moves
		for _, move := range moves {
			nextPos := curr.Pos + move

			// Allow staying at initial position or moving into/through valley
			if nextPos == initialPos || nextPos == end+1i || v.isValid(nextPos, nextTime) {
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
