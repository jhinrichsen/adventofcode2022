package adventofcode2022

type Node struct {
	Value         int
	OriginalIndex int
	Prev          *Node
	Next          *Node
}

func Day20(srcs []int, mix int, part1 bool) int {
	// Apply decryption key for part 2
	const decryptionKey = 811589153
	values := make([]int, len(srcs))
	for i := range srcs {
		if part1 {
			values[i] = srcs[i]
		} else {
			values[i] = srcs[i] * decryptionKey
		}
	}

	n := len(values)
	// Create circular doubly-linked list
	nodes := make([]*Node, n)
	for i := range values {
		nodes[i] = &Node{Value: values[i], OriginalIndex: i}
	}

	// Link nodes in a circle
	for i := range nodes {
		nodes[i].Prev = nodes[(i-1+n)%n]
		nodes[i].Next = nodes[(i+1)%n]
	}

	// Mix the specified number of times
	for range mix {
		// Process each number in original order
		for i := range nodes {
			node := nodes[i]
			if node.Value == 0 {
				continue
			}

			// Remove node from current position
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev

			// Move to new position
			steps := node.Value % (n - 1)
			if steps < 0 {
				steps += (n - 1)
			}

			// Find insertion point
			target := node.Prev
			for range steps {
				target = target.Next
			}

			// Insert after target
			node.Next = target.Next
			node.Prev = target
			target.Next.Prev = node
			target.Next = node
		}
	}

	// Find node with value 0
	var zeroNode *Node
	for i := range nodes {
		if nodes[i].Value == 0 {
			zeroNode = nodes[i]
			break
		}
	}

	// Get values at 1000, 2000, 3000 after zero
	current := zeroNode
	for range 1000 {
		current = current.Next
	}
	v1 := current.Value

	for range 1000 {
		current = current.Next
	}
	v2 := current.Value

	for range 1000 {
		current = current.Next
	}
	v3 := current.Value

	return v1 + v2 + v3
}
