package adventofcode2022

import (
	"encoding/json"
	"slices"
)

// Day13 solves day 13: Distress Signal
func Day13(lines []string, part1 bool) uint {
	if part1 {
		return day13Part1(lines)
	}
	return day13Part2(lines)
}

func day13Part1(lines []string) uint {
	var sum uint
	pairIdx := 1

	for i := 0; i+1 < len(lines); i += 3 {
		var left, right any
		err := json.Unmarshal([]byte(lines[i]), &left)
		if err != nil {
			pairIdx++
			continue
		}

		err = json.Unmarshal([]byte(lines[i+1]), &right)
		if err != nil {
			pairIdx++
			continue
		}

		if comparePackets(left, right) < 0 {
			sum += uint(pairIdx)
		}
		pairIdx++
	}

	return sum
}

func day13Part2(lines []string) uint {
	// Parse all packets (skip blank lines)
	var packets []any

	for _, line := range lines {
		if line == "" {
			continue
		}
		var packet any
		err := json.Unmarshal([]byte(line), &packet)
		if err != nil {
			continue
		}
		packets = append(packets, packet)
	}

	// Add divider packets
	var divider1, divider2 any
	err := json.Unmarshal([]byte("[[2]]"), &divider1)
	if err != nil {
		return 0
	}
	err = json.Unmarshal([]byte("[[6]]"), &divider2)
	if err != nil {
		return 0
	}
	packets = append(packets, divider1, divider2)

	// Sort packets using comparePackets
	slices.SortFunc(packets, comparePackets)

	// Find indices of divider packets (1-indexed)
	var idx1, idx2 uint
	for i, packet := range packets {
		if comparePackets(packet, divider1) == 0 {
			idx1 = uint(i + 1)
		}
		if comparePackets(packet, divider2) == 0 {
			idx2 = uint(i + 1)
		}
	}

	return idx1 * idx2
}

// comparePackets compares two packets according to the rules:
// Returns -1 if left < right (correct order)
// Returns 0 if left == right (continue checking)
// Returns 1 if left > right (wrong order)
func comparePackets(left, right any) int {
	// Both are numbers
	lNum, lIsNum := left.(float64)
	rNum, rIsNum := right.(float64)

	if lIsNum && rIsNum {
		if lNum < rNum {
			return -1
		}
		if lNum > rNum {
			return 1
		}
		return 0
	}

	// Convert to slices if needed
	lSlice, lIsSlice := left.([]any)
	rSlice, rIsSlice := right.([]any)

	if !lIsSlice {
		lSlice = []any{left}
	}
	if !rIsSlice {
		rSlice = []any{right}
	}

	// Compare lists element by element
	for i := 0; i < len(lSlice) && i < len(rSlice); i++ {
		cmp := comparePackets(lSlice[i], rSlice[i])
		if cmp != 0 {
			return cmp
		}
	}

	// If all elements are equal, compare lengths
	if len(lSlice) < len(rSlice) {
		return -1
	}
	if len(lSlice) > len(rSlice) {
		return 1
	}
	return 0
}
