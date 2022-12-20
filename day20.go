package adventofcode2022

import "fmt"

func Day20(srcs []int, mix int, part1 bool) int {
	var (
		l    = len(srcs)
		dsts = make([]int, l)

		ring = func(i int) int {
			return i % l
		}

		dst = func(i int) int {
			return dsts[ring(i)]
		}

		findIndex = func(n int) int {
			for i := 0; i < l; i++ {
				if dsts[i] == n {
					return i
				}
			}
			return -1
		}
	)

	copy(dsts, srcs)
	fmt.Printf("initial arrangement:\n")
	fmt.Printf("%v\n", dsts)
	for j := 0; j < mix; j++ {
		for k := 0; k < l; k++ {
			var (
				rel       = srcs[k]
				fromIndex = findIndex(rel)
				intoIndex int
				count     int
				interim   = make([]int, l)
				add       = func(is []int) {
					count += copy(interim[count:], is)
				}
			)
			if rel > 0 {
				intoIndex = ring(fromIndex + rel)
			} else {
				intoIndex = ring(l + rel)
			}

			add(dsts[0:fromIndex])               // copy up to the remove position
			add(dsts[fromIndex+1 : intoIndex+1]) // skip remove position
			add(srcs[k : k+1])                   // add original value
			add(dsts[intoIndex+2:])              // add the rest
			copy(dsts, interim)

			fmt.Printf("%d moves between %d and %d:\n", rel,
				dsts[findIndex(dsts[intoIndex])-1],
				dsts[findIndex(dsts[intoIndex])-1])
			fmt.Printf("%v\n", dsts)
		}
	}

	j := findIndex(0)
	d1 := dst(j + 1000 + 1)
	d2 := dst(j + 2000 + 1)
	d3 := dst(j + 3000 + 1)
	return d1 + d2 + d3
}
