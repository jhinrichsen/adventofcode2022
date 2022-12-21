package adventofcode2022

func Day20(srcs []int, mix int, part1 bool) int {
	var (
		dim  = len(srcs)
		dsts = make([]int, dim)

		ring = func(i int) int {
			// i % dim returns negative remainder for negative i
			return (i%dim + dim) % dim
		}

		findIndex = func(n int) int {
			for i := 0; i < dim; i++ {
				if dsts[i] == n {
					return i
				}
			}
			return -1
		}

		dst = func(i int) int {
			return dsts[ring(i)]
		}
	)

	copy(dsts, srcs)
	for j := 0; j < mix; j++ {
		for k := 0; k < dim; k++ {
			var (
				val       = srcs[k]
				rel       = val
				fromIndex = findIndex(val)
				step      = 1
			)

			if rel < 0 {
				rel = -rel
				step = -1
			}
			if step < 0 && fromIndex-rel <= 0 {
				rel = dim - (rel % dim) - 1
				step = 1
			} else if step > 0 && fromIndex+rel >= dim {
				rel = dim - (rel % dim) - 1
				step = -1
			}
			from := ring(fromIndex)
			into := ring(from + step)
			for l := 0; l < rel; l++ {
				dsts[from], dsts[into] = dsts[into], dsts[from]
				from = ring(from + step)
				into = ring(into + step)
			}
		}
	}

	j := findIndex(0)
	d1 := dst(j + 1000)
	d2 := dst(j + 2000)
	d3 := dst(j + 3000)
	return d1 + d2 + d3
}
