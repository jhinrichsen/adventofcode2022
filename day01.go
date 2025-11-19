package adventofcode2022

import (
	"slices"
	"strconv"
)

// Day01 finds the sum of calories carried by the top elf (part1) or top 3 elves (part2).
func Day01(lines []string, part1 bool) uint {
	// Ensure final empty line to trigger append of last elf
	if len(lines) > 0 && lines[len(lines)-1] != "" {
		lines = append(lines, "")
	}

	var elves []uint
	var current uint

	for _, line := range lines {
		if line == "" {
			if current > 0 {
				elves = append(elves, current)
				current = 0
			}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			continue // Skip invalid lines
		}
		current += uint(calories)
	}

	// Sort descending
	slices.Sort(elves)
	slices.Reverse(elves)

	n := 3
	if part1 {
		n = 1
	}

	var sum uint
	for i := range min(n, len(elves)) {
		sum += elves[i]
	}

	return sum
}
