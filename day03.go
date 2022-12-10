package adventofcode2022

import (
	"strings"
)

func Day03(lines []string, part1 bool) int {
	// find character that appear in both left and right
	intersect := func(left, right string, n int) string {
		var sb strings.Builder
		var count int
		for i := 0; i < len(left); i++ {
			// find b in right
			for j := 0; j < len(right); j++ {
				if left[i] == right[j] {
					sb.WriteByte(left[i])
					count++
					if count == n {
						return sb.String()
					}
				}
			}
		}
		return sb.String()
	}
	prio := func(b byte) byte {
		// A..Z = 65..90 -> 27..
		n := b - 38
		// a = 97..122 -> 1..
		if n > 52 {
			n -= 58
		}
		return n
	}

	sum := 0
	if part1 {
		// each line is two compartments
		for _, rucksack := range lines {
			items := len(rucksack)
			compartment := items / 2
			left := rucksack[0:compartment]
			right := rucksack[compartment:items]
			sum += int(prio(intersect(left, right, 1)[0]))
		}
	} else {
		// part 2: group of three
		for i := 0; i < len(lines); i += 3 {
			// Haskell teaches you that there is no function that
			// has more than one parameter
			s := intersect(lines[i], lines[i+1], -1)
			s = intersect(s, lines[i+2], 1)
			sum += int(prio(s[0]))
		}
	}
	return sum
}
