package adventofcode2022

// Day01 finds the sum of calories carried by the top elf (part1) or top 3 elves (part2).
func Day01(lines []string, part1 bool) uint {
	var top1, top2, top3 uint
	var current uint

	for _, line := range lines {
		if line == "" {
			if current > 0 {
				// Insert into top 3
				if current > top1 {
					top3 = top2
					top2 = top1
					top1 = current
				} else if current > top2 {
					top3 = top2
					top2 = current
				} else if current > top3 {
					top3 = current
				}
				current = 0
			}
			continue
		}

		// Manual parsing - faster than strconv.Atoi
		var n uint
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= '0' && c <= '9' {
				n = n*10 + uint(c-'0')
			}
		}
		current += n
	}

	// Capture final elf
	if current > 0 {
		if current > top1 {
			top3 = top2
			top2 = top1
			top1 = current
		} else if current > top2 {
			top3 = top2
			top2 = current
		} else if current > top3 {
			top3 = current
		}
	}

	if part1 {
		return top1
	}
	return top1 + top2 + top3
}
