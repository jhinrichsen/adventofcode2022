package adventofcode2022

import (
	"strconv"
	"strings"
)

// Day04 counts pairs with fully contained ranges (part1) or overlapping ranges (part2).
func Day04(lines []string, part1 bool) int {
	f := Overlaps
	if part1 {
		f = Contains
	}

	count := 0
	for _, line := range lines {
		x1, x2, x3, x4, ok := parse(line)
		if !ok {
			continue // Skip invalid lines
		}
		if f(x1, x2, x3, x4) {
			count++
		}
	}
	return count
}

func parse(line string) (int, int, int, int, bool) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return 0, 0, 0, 0, false
	}

	l1 := strings.Split(parts[0], "-")
	l2 := strings.Split(parts[1], "-")
	if len(l1) != 2 || len(l2) != 2 {
		return 0, 0, 0, 0, false
	}

	ss := []string{l1[0], l1[1], l2[0], l2[1]}
	var ns [4]int
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, 0, 0, 0, false
		}
		ns[i] = n
	}
	return ns[0], ns[1], ns[2], ns[3], true
}

// Contains returns true if range a [a1..a2] is within b [b1..b2] or vice versa.
func Contains(a1, a2, b1, b2 int) bool {
	return a1 >= b1 && a2 <= b2 || // a within b
		a1 <= b1 && a2 >= b2 // b within a
}

// Overlaps returns true if range a [a1..a2] overlaps b [b1..b2] or vice versa.
func Overlaps(a1, a2, b1, b2 int) bool {
	return Contains(a1, a2, b1, b2) ||
		a1 >= b1 && a1 <= b2 || // a1 in b
		a2 >= b1 && a2 <= b2 // a2 in b
}
