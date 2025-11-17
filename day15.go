package adventofcode2022

import (
	"image"
	"strconv"
	"strings"
)

// Day15 solves day 15: Beacon Exclusion Zone
func Day15(lines []string, targetY int, part1 bool) uint {
	sensors := parseSensors(lines)

	if part1 {
		return day15Part1(sensors, targetY)
	}
	return 0 // Part 2 placeholder
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

func mergeIntervals(intervals [][2]int) [][2]int {
	if len(intervals) == 0 {
		return nil
	}

	// Sort intervals by start position
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] < intervals[i][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

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
