package adventofcode2022

import (
	"slices"
	"strconv"
)

// Day01 finds the sum of calories carried by the top elf (part1) or top 3 elves (part2).
func Day01(lines []string, part1 bool) uint {
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

	// Capture final elf if not followed by empty line
	if current > 0 {
		elves = append(elves, current)
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
