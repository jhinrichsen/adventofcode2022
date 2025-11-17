package adventofcode2022

import (
	"image"
	"slices"
	"strconv"
	"strings"
)

// Day15 solves day 15: Beacon Exclusion Zone
func Day15(lines []string, targetY, maxCoord int, part1 bool) uint {
	sensors := parseSensors(lines)

	if part1 {
		return day15Part1(sensors, targetY)
	}
	return day15Part2(sensors, maxCoord)
}

type Sensor struct {
	pos    image.Point
	beacon image.Point
	radius int
}

func parseSensors(lines []string) []Sensor {
	var sensors []Sensor

	for _, line := range lines {
		// Parse: "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}

		sensorPart := strings.TrimPrefix(parts[0], "Sensor at ")
		beaconPart := strings.TrimPrefix(parts[1], "closest beacon is at ")

		sensorPos := parseCoord(sensorPart)
		beaconPos := parseCoord(beaconPart)

		radius := manhattan(sensorPos, beaconPos)
		sensors = append(sensors, Sensor{sensorPos, beaconPos, radius})
	}

	return sensors
}

func parseCoord(s string) image.Point {
	// Parse "x=2, y=18"
	parts := strings.Split(s, ", ")
	if len(parts) != 2 {
		return image.Point{}
	}

	xStr := strings.TrimPrefix(parts[0], "x=")
	yStr := strings.TrimPrefix(parts[1], "y=")

	x, err := strconv.Atoi(xStr)
	if err != nil {
		return image.Point{}
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		return image.Point{}
	}

	return image.Point{X: x, Y: y}
}

func manhattan(p1, p2 image.Point) int {
	dx := p1.X - p2.X
	if dx < 0 {
		dx = -dx
	}
	dy := p1.Y - p2.Y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func day15Part1(sensors []Sensor, targetY int) uint {
	// Find all intervals covered in targetY row
	var intervals [][2]int

	for _, s := range sensors {
		// Check if sensor's exclusion zone reaches targetY
		dy := s.pos.Y - targetY
		if dy < 0 {
			dy = -dy
		}
		if dy > s.radius {
			continue // This sensor doesn't reach targetY
		}

		// Calculate the X range covered at targetY
		remaining := s.radius - dy
		x1 := s.pos.X - remaining
		x2 := s.pos.X + remaining

		intervals = append(intervals, [2]int{x1, x2})
	}

	// Merge overlapping intervals
	merged := mergeIntervals(intervals)

	// Count total positions covered
	var count uint
	for _, interval := range merged {
		count += uint(interval[1] - interval[0] + 1)
	}

	// Subtract beacon positions in targetY
	beacons := make(map[int]bool)
	for _, s := range sensors {
		if s.beacon.Y == targetY {
			beacons[s.beacon.X] = true
		}
	}

	// Check which beacons are in the covered range
	for bx := range beacons {
		for _, interval := range merged {
			if bx >= interval[0] && bx <= interval[1] {
				count--
				break
			}
		}
	}

	return count
}

func day15Part2(sensors []Sensor, maxCoord int) uint {
	// Pre-allocate intervals slice and reuse it
	intervals := make([][2]int, 0, len(sensors))

	// Search for the one uncovered position in [0, maxCoord] x [0, maxCoord]
	for y := 0; y <= maxCoord; y++ {
		// Reuse intervals slice
		intervals = intervals[:0]

		for i := range sensors {
			s := &sensors[i]
			// Check if sensor's exclusion zone reaches y
			dy := s.pos.Y - y
			if dy < 0 {
				dy = -dy
			}
			if dy > s.radius {
				continue
			}

			// Calculate the X range covered at y
			remaining := s.radius - dy
			x1 := s.pos.X - remaining
			x2 := s.pos.X + remaining

			// Clamp to search bounds
			if x1 < 0 {
				x1 = 0
			}
			if x2 > maxCoord {
				x2 = maxCoord
			}

			intervals = append(intervals, [2]int{x1, x2})
		}

		// Merge intervals
		merged := mergeIntervalsFast(intervals)

		// Check for gap in coverage
		if len(merged) == 0 {
			// Entire row is uncovered (shouldn't happen)
			return uint(y)
		}

		// Check if there's a gap in [0, maxCoord]
		if merged[0][0] > 0 {
			// Gap at the start
			return uint(y)
		}

		// Check for gap between merged intervals
		for i := 0; i < len(merged)-1; i++ {
			if merged[i][1]+1 < merged[i+1][0] {
				// Found a gap
				x := merged[i][1] + 1
				return uint(x)*4000000 + uint(y)
			}
		}

		// Check if coverage ends before maxCoord
		if merged[len(merged)-1][1] < maxCoord {
			x := merged[len(merged)-1][1] + 1
			return uint(x)*4000000 + uint(y)
		}
	}

	return 0
}

func mergeIntervals(intervals [][2]int) [][2]int {
	if len(intervals) == 0 {
		return nil
	}

	// Sort intervals by start position using slices.SortFunc
	slices.SortFunc(intervals, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	var merged [][2]int
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= current[1]+1 {
			// Overlapping or adjacent, merge them
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			// No overlap, save current and start new
			merged = append(merged, current)
			current = intervals[i]
		}
	}
	merged = append(merged, current)

	return merged
}

func mergeIntervalsFast(intervals [][2]int) [][2]int {
	if len(intervals) == 0 {
		return nil
	}

	// Sort intervals by start position using slices.SortFunc
	slices.SortFunc(intervals, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	// Merge in-place to avoid allocations
	writeIdx := 0
	for readIdx := 1; readIdx < len(intervals); readIdx++ {
		if intervals[readIdx][0] <= intervals[writeIdx][1]+1 {
			// Overlapping or adjacent, merge
			if intervals[readIdx][1] > intervals[writeIdx][1] {
				intervals[writeIdx][1] = intervals[readIdx][1]
			}
		} else {
			// No overlap, move to next position
			writeIdx++
			intervals[writeIdx] = intervals[readIdx]
		}
	}

	return intervals[:writeIdx+1]
}
