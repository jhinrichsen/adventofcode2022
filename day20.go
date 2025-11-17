package adventofcode2022

type Node struct {
	Value         int
	OriginalIndex int
}

func Day20(srcs []int, mix int, part1 bool) int {
	// Create list with original indices
	list := make([]Node, len(srcs))
	for i := range srcs {
		list[i] = Node{Value: srcs[i], OriginalIndex: i}
	}

	// Mix the specified number of times
	for range mix {
		// Process each number in original order
		for origIdx := range srcs {
			// Find current position of this number
			currIdx := -1
			for i := range list {
				if list[i].OriginalIndex == origIdx {
					currIdx = i
					break
				}
			}

			// Remove the node
			node := list[currIdx]
			list = append(list[:currIdx], list[currIdx+1:]...)

			// Calculate new position (modulo list length after removal)
			newIdx := currIdx + node.Value
			if len(list) > 0 {
				newIdx = ((newIdx % len(list)) + len(list)) % len(list)
			}

			// Insert at new position
			list = append(list[:newIdx], append([]Node{node}, list[newIdx:]...)...)
		}
	}

	// Find index of 0
	zeroIdx := -1
	for i := range list {
		if list[i].Value == 0 {
			zeroIdx = i
			break
		}
	}

	// Get values at 1000, 2000, 3000 after zero
	v1 := list[(zeroIdx+1000)%len(list)].Value
	v2 := list[(zeroIdx+2000)%len(list)].Value
	v3 := list[(zeroIdx+3000)%len(list)].Value

	return v1 + v2 + v3
}
