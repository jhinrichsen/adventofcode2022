package adventofcode2021

func Day03(lines []string) int {
	sum := 0
	for _, rucksack := range lines {
		items := len(rucksack)
		compartment := items / 2
		left := []byte(rucksack[0:compartment])
		right := []byte(rucksack[compartment:items])
		for i := 0; i < len(left); i++ {
			// find b in right
			for j := 0; j < len(right); j++ {
				if left[i] == right[j] {
					// A..Z = 65..90 -> 27..
					priority := left[i] - 38
					// a = 97..122 -> 1..
					if priority > 52 {
						priority -= 58
					}
					sum += int(priority)

					// short circuit: one match per rucksack
					i, j = items, items
				}
			}
		}
	}
	return sum
}
