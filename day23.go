package adventofcode2022

import "image"

type point = image.Point

var (
	north     = point{0, -1}
	south     = point{0, 1}
	west      = point{-1, 0}
	east      = point{1, 0}
	northEast = point{1, -1}
	northWest = point{-1, -1}
	southEast = point{1, 1}
	southWest = point{-1, 1}
)

func Day23(lines []string, part1 bool) uint {
	// Parse into set of elf positions
	elvesMap := make(map[point]struct{}, 4000)
	elvesSlice := make([]point, 0, 4000)
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '#' {
				p := point{x, y}
				elvesMap[p] = struct{}{}
				elvesSlice = append(elvesSlice, p)
			}
		}
	}

	proposals := [4]struct {
		checks [3]point
		move   point
	}{
		{[3]point{north, northEast, northWest}, north},
		{[3]point{south, southEast, southWest}, south},
		{[3]point{west, northWest, southWest}, west},
		{[3]point{east, northEast, southEast}, east},
	}

	neighbors := [8]point{north, northEast, east, southEast, south, southWest, west, northWest}

	maxRounds := 10
	if !part1 {
		maxRounds = 1000000
	}

	// Reuse maps to reduce allocations
	dsts := make(map[point]point, len(elvesSlice))
	counts := make(map[point]int, len(elvesSlice))
	proposalOffset := 0

	for round := 1; round <= maxRounds; round++ {
		// Clear maps for reuse
		clear(dsts)
		clear(counts)

		// First half: propose moves - iterate over slice for better cache locality
		for i := range elvesSlice {
			elf := elvesSlice[i]

			// Check if elf has any neighbors
			hasNeighbor := false
			for _, dir := range neighbors {
				if _, exists := elvesMap[elf.Add(dir)]; exists {
					hasNeighbor = true
					break
				}
			}

			if !hasNeighbor {
				continue
			}

			// Try each proposal direction (use offset instead of rotating slice)
			for pi := 0; pi < 4; pi++ {
				prop := proposals[(proposalOffset+pi)&3]
				canMove := true
				for _, check := range prop.checks {
					if _, exists := elvesMap[elf.Add(check)]; exists {
						canMove = false
						break
					}
				}

				if canMove {
					dst := elf.Add(prop.move)
					dsts[elf] = dst
					counts[dst]++
					break
				}
			}
		}

		// Second half: execute moves
		stable := len(dsts) == 0
		if stable && !part1 {
			return uint(round)
		}

		if !stable {
			// Execute moves
			for from, to := range dsts {
				if counts[to] == 1 {
					delete(elvesMap, from)
					elvesMap[to] = struct{}{}
				}
			}
			// Rebuild elvesSlice from updated elvesMap
			elvesSlice = elvesSlice[:0]
			for elf := range elvesMap {
				elvesSlice = append(elvesSlice, elf)
			}
		}

		// Rotate proposals using offset
		proposalOffset = (proposalOffset + 1) & 3
	}

	// Count empty tiles in bounding rectangle
	minX, maxX, minY, maxY := 1<<30, -1<<30, 1<<30, -1<<30
	for p := range elvesMap {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	area := (maxX - minX + 1) * (maxY - minY + 1)
	return uint(area - len(elvesMap))
}
