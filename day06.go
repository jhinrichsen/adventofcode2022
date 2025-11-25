package adventofcode2022

import (
	"bytes"
	"math/bits"
)

// Day06 finds the position of the first start-of-packet (part1, size=4) or
// start-of-message (part2, size=14) marker in the datastream.
func Day06(input []byte, part1 bool) (int, error) {
	s := string(bytes.TrimSpace(input))

	size := 14
	if part1 {
		size = 4
	}

	// Optimized bitset approach using OnesCount
	for i := 0; i < len(s)-size; i++ {
		var marker uint32
		for j := i; j < i+size; j++ {
			marker |= 1 << (s[j] - 'a')
		}
		if bits.OnesCount32(marker) == size {
			return i + size, nil
		}
	}
	return 0, nil
}
