package adventofcode2022

func Day10(lines []string) int {
	dim := 2 * len(lines) // 2 states/ op max
	states := make([]int, dim)

	cycle := 1 // 1-based
	x := 1
	for _, line := range lines {
		if line == "noop" {
			states[cycle] = x
			cycle++
		} else {
			// cycle 1
			states[cycle] = x
			cycle++

			// cycle 2
			states[cycle] = x
			// value changes _after_ cycle
			cycle++

			// roll our own parser, because why not?
			/*
				n, _ := strconv.Atoi(line[5:])
				x += n
			*/

			n := 0
			sign := 1
			i := 5
			if line[i] == '-' {
				sign = -1
				i++
			}
			for ; i < len(line); i++ {
				n = 10*n + int(line[i]-'0')
			}
			x += sign * n
		}
	}

	sum := 0
	for cycle := 20; cycle <= 220; cycle += 40 {
		sum += cycle * states[cycle]
	}
	return sum
}
