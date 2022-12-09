package adventofcode2021

func Day06(s string) int {
	const size = 4
	window := make(map[byte]int)
	add := func(c byte) {
		window[c] += 1
	}
	del := func(c byte) {
		n := window[c]
		if n > 1 {
			window[c] -= 1
		} else {
			delete(window, c)
		}
	}
	hasMarker := func() bool {
		return len(window) == size
	}

	// populate window
	for i := 0; i < size; i++ {
		add(s[i])
	}
	if hasMarker() {
		return size
	}

	// slide window through stream
	for i := size; i < len(s); i++ {
		del(s[i-size])
		add(s[i])
		if hasMarker() {
			return i + 1 // 1-based position
		}
	}
	return 0
}
