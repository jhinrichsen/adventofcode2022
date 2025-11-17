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
	elves := make(map[point]struct{}, 4000)
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '#' {
				elves[point{x, y}] = struct{}{}
			}
		}
	}

	proposals := []struct {
		checks []point
		move   point
	}{
		{[]point{north, northEast, northWest}, north},
		{[]point{south, southEast, southWest}, south},
		{[]point{west, northWest, southWest}, west},
		{[]point{east, northEast, southEast}, east},
	}

	neighbors := [8]point{north, northEast, east, southEast, south, southWest, west, northWest}

	maxRounds := 10
	if !part1 {
		maxRounds = 1000000
	}

	// Reuse maps to reduce allocations
	dsts := make(map[point]point, len(elves))
	counts := make(map[point]int, len(elves))

	for round := 1; round <= maxRounds; round++ {
		// Clear maps for reuse
		clear(dsts)
		clear(counts)

		// First half: propose moves
		for elf := range elves {
			// Check if elf has any neighbors
			hasNeighbor := false
			for _, dir := range neighbors {
				if _, exists := elves[elf.Add(dir)]; exists {
					hasNeighbor = true
					break
				}
			}

			if !hasNeighbor {
				continue
			}

			// Try each proposal direction
			for _, prop := range proposals {
				canMove := true
				for _, check := range prop.checks {
					if _, exists := elves[elf.Add(check)]; exists {
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
			for from, to := range dsts {
				if counts[to] == 1 {
					delete(elves, from)
					elves[to] = struct{}{}
				}
			}
		}

		// Rotate proposals
		prop0 := proposals[0]
		copy(proposals, proposals[1:])
		proposals[len(proposals)-1] = prop0
	}

	// Count empty tiles in bounding rectangle
	minX, maxX, minY, maxY := 1<<30, -1<<30, 1<<30, -1<<30
	for p := range elves {
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
	return uint(area - len(elves))
}
