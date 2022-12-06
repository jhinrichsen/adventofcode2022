package adventofcode2021

import (
	"sort"
	"strconv"
)

// Day01 is the first puzzle.
// It returns the most calories of any raindeer.
// In case of an empty list 0 is returned.
// Non-numbers are ignored.
// No error checking for negative calories.
func Day01(lines []string, n int) int {
	// make sure we have one final newline
	if lines[len(lines)-1] != "" {
		lines = append(lines, "")
	}

	var calories []int // calories per raindeer
	var sum int
	for _, line := range lines {
		if line == "" {
			calories = append(calories, sum)
			sum = 0
		} else {
			n, _ := strconv.Atoi(line)
			sum += n
		}
	}

	// go for the more efficient custom order instead of Reverse() Sort()
	sort.Slice(calories, func(a, b int) bool {
		return calories[a] > calories[b]
	})

	// first n largest calories
	var total int
	for i := 0; i < n; i++ {
		total += calories[i]
	}
	return total
}
