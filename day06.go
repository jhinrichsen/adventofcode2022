package adventofcode2021

import (
	"fmt"
	"math/bits"
)

func Day06(s string) int {
	fmt.Printf("testing %s\n", s)
	const size = 4
	// prerequisite: string consists of 26 lower case ASCII characters [a..z]
	var window uint32
	hasMarker := func() bool {
		return bits.OnesCount32(window) >= size
	}
	// preload marker window
	for i := 0; i < size; i++ {
		window |= 1 << (s[i] - 'a')
		fmt.Printf("pre %c: %32b\n", s[i], window)
	}
	if hasMarker() {
		return size
	}

	// mill-cut through stream in a 4 byte window
	for i := size; i < len(s); i++ {
		window &= ^(1 << (s[i-size] - 'a')) // drop oldest one
		fmt.Printf("%2d. rem %c: %32b\n", i, s[i-size], window)
		window |= 1 << (s[i] - 'a') // add current
		fmt.Printf("%2d. add %c: %32b (%d)\n", i, s[i], window, bits.OnesCount32(window))
		if hasMarker() {
			return i + 1 // result must be 1-based
		}
	}
	return 0
}
