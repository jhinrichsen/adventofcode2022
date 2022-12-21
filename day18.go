package adventofcode2022

import (
	"math"
	"strconv"
	"strings"
)

func Day18(lines []string, part1 bool) int {
	m := make(map[D3]bool)
	neighbors := func(k D3) int {
		prospects := []D3{
			{k.x - 1, k.y, k.z},
			{k.x + 1, k.y, k.z},
			{k.x, k.y - 1, k.z},
			{k.x, k.y + 1, k.z},
			{k.x, k.y, k.z - 1},
			{k.x, k.y, k.z + 1},
		}
		var filled int
		for i := range prospects {
			b := m[prospects[i]]
			if b {
				filled++
			}
		}
		return filled
	}
	isHole := func(k D3) bool {
		return !m[k]
	}
	isSurrounded := func(k D3) bool {
		return neighbors(k) == 6
	}

	for _, line := range lines {
		ss := strings.Split(line, ",")
		x, _ := strconv.Atoi(ss[0])
		y, _ := strconv.Atoi(ss[1])
		z, _ := strconv.Atoi(ss[2])
		m[D3{x, y, z}] = true
	}

	surface := 0
	for k := range m {
		surface += 6 - neighbors(k)
	}
	if part1 {
		return surface
	}

	// calculate dimensions
	var x0, y0, z0 int = math.MaxInt, math.MaxInt, math.MaxInt
	var x1, y1, z1 int = math.MinInt, math.MinInt, math.MinInt
	for k := range m {
		if k.x < x0 {
			x0 = k.x
		}
		if k.x > x1 {
			x1 = k.x
		}
		if k.y < y0 {
			y0 = k.y
		}
		if k.y > y1 {
			y1 = k.y
		}
		if k.z < z0 {
			z0 = k.z
		}
		if k.z > z1 {
			z1 = k.z
		}
	}
	// fmt.Printf("found area from %d,%d,%d to %d,%d,%d\n", x0, y0, z0, x1, y1, z1)
	var airTraps int
	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			for z := z0; z < z1; z++ {
				k := D3{x, y, z}
				if isHole(k) {
					// fmt.Printf("found air trap at %d,%d,%d with %d neighbors\n", x, y, z, neighbors(k))
					if isSurrounded(k) {
						// fmt.Printf("found surrounded trap at %d,%d,%d with %d neighbors\n", x, y, z, neighbors(k))
						airTraps++
					}
				}
			}
		}
	}
	return surface - 6*airTraps
}

type D3 struct {
	x, y, z int
}
