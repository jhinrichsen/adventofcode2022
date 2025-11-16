package adventofcode2022

import (
	"image"
	"strconv"
	"strings"
)

// Day14 solves day 14: Regolith Reservoir
func Day14(lines []string, part1 bool) uint {
	// Parse rock structures
	rocks := make(map[image.Point]bool)
	maxY := 0

	for _, line := range lines {
		points := parseRockLine(line)
		for i := 1; i < len(points); i++ {
			drawLine(rocks, points[i-1], points[i])
			if points[i].Y > maxY {
				maxY = points[i].Y
			}
			if points[i-1].Y > maxY {
				maxY = points[i-1].Y
			}
		}
	}

	// Simulate sand falling
	var count uint
	source := image.Point{X: 500, Y: 0}

	for {
		if part1 {
			if !dropSand(rocks, source, maxY) {
				break
			}
		} else {
			if !dropSandPart2(rocks, source, maxY+2) {
				break
			}
		}
		count++
	}

	return count
}

// parseRockLine parses a line like "498,4 -> 498,6 -> 496,6" into points
func parseRockLine(line string) []image.Point {
	var points []image.Point
	coords := strings.Split(line, " -> ")

	for _, coord := range coords {
		parts := strings.Split(coord, ",")
		if len(parts) != 2 {
			continue
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		points = append(points, image.Point{X: x, Y: y})
	}

	return points
}

// drawLine draws a rock line between two points
func drawLine(rocks map[image.Point]bool, p1, p2 image.Point) {
	if p1.X == p2.X {
		// Vertical line
		startY, endY := p1.Y, p2.Y
		if startY > endY {
			startY, endY = endY, startY
		}
		for y := startY; y <= endY; y++ {
			rocks[image.Point{X: p1.X, Y: y}] = true
		}
	} else {
		// Horizontal line
		startX, endX := p1.X, p2.X
		if startX > endX {
			startX, endX = endX, startX
		}
		for x := startX; x <= endX; x++ {
			rocks[image.Point{X: x, Y: p1.Y}] = true
		}
	}
}

// dropSand simulates one unit of sand falling (part 1)
// Returns true if sand came to rest, false if it fell into abyss
func dropSand(rocks map[image.Point]bool, source image.Point, maxY int) bool {
	sand := source

	for {
		// Check if sand fell into abyss
		if sand.Y > maxY {
			return false
		}

		// Try to fall down
		down := image.Point{X: sand.X, Y: sand.Y + 1}
		if !rocks[down] {
			sand = down
			continue
		}

		// Try to fall down-left
		downLeft := image.Point{X: sand.X - 1, Y: sand.Y + 1}
		if !rocks[downLeft] {
			sand = downLeft
			continue
		}

		// Try to fall down-right
		downRight := image.Point{X: sand.X + 1, Y: sand.Y + 1}
		if !rocks[downRight] {
			sand = downRight
			continue
		}

		// Sand comes to rest
		rocks[sand] = true
		return true
	}
}

// dropSandPart2 simulates one unit of sand falling (part 2 with floor)
// Returns true if sand came to rest (not at source), false if source is blocked
func dropSandPart2(rocks map[image.Point]bool, source image.Point, floorY int) bool {
	sand := source

	for {
		// Check if sand reached floor
		if sand.Y+1 == floorY {
			rocks[sand] = true
			return sand != source
		}

		// Try to fall down
		down := image.Point{X: sand.X, Y: sand.Y + 1}
		if !rocks[down] {
			sand = down
			continue
		}

		// Try to fall down-left
		downLeft := image.Point{X: sand.X - 1, Y: sand.Y + 1}
		if !rocks[downLeft] {
			sand = downLeft
			continue
		}

		// Try to fall down-right
		downRight := image.Point{X: sand.X + 1, Y: sand.Y + 1}
		if !rocks[downRight] {
			sand = downRight
			continue
		}

		// Sand comes to rest
		rocks[sand] = true
		return sand != source
	}
}
