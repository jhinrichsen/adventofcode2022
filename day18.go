package adventofcode2022

import (
	"math"
	"strconv"
	"strings"
)

func Day18(lines []string, part1 bool) uint {
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

	for _, line := range lines {
		ss := strings.Split(line, ",")
		if len(ss) != 3 {
			continue
		}
		x, err := strconv.Atoi(ss[0])
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(ss[1])
		if err != nil {
			continue
		}
		z, err := strconv.Atoi(ss[2])
		if err != nil {
			continue
		}
		m[D3{x, y, z}] = true
	}

	surface := 0
	for k := range m {
		surface += 6 - neighbors(k)
	}
	if part1 {
		return uint(surface)
	}

	// Part 2: Find exterior surface area only
	// Calculate bounding box and expand by 1
	var x0, y0, z0 int = math.MaxInt, math.MaxInt, math.MaxInt
	var x1, y1, z1 int = math.MinInt, math.MinInt, math.MinInt
	for k := range m {
		x0 = min(x0, k.x)
		x1 = max(x1, k.x)
		y0 = min(y0, k.y)
		y1 = max(y1, k.y)
		z0 = min(z0, k.z)
		z1 = max(z1, k.z)
	}
	// Expand bounding box by 1 to ensure we start outside
	x0--
	y0--
	z0--
	x1++
	y1++
	z1++

	// Flood fill from outside to find all exterior air cubes
	exterior := make(map[D3]bool)
	stack := []D3{{x0, y0, z0}}
	exterior[D3{x0, y0, z0}] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		prospects := []D3{
			{current.x - 1, current.y, current.z},
			{current.x + 1, current.y, current.z},
			{current.x, current.y - 1, current.z},
			{current.x, current.y + 1, current.z},
			{current.x, current.y, current.z - 1},
			{current.x, current.y, current.z + 1},
		}

		for _, p := range prospects {
			// Skip if out of bounds
			if p.x < x0 || p.x > x1 || p.y < y0 || p.y > y1 || p.z < z0 || p.z > z1 {
				continue
			}
			// Skip if it's lava
			if m[p] {
				continue
			}
			// Skip if already visited
			if exterior[p] {
				continue
			}
			exterior[p] = true
			stack = append(stack, p)
		}
	}

	// Count lava faces adjacent to exterior air
	exteriorSurface := 0
	for k := range m {
		prospects := []D3{
			{k.x - 1, k.y, k.z},
			{k.x + 1, k.y, k.z},
			{k.x, k.y - 1, k.z},
			{k.x, k.y + 1, k.z},
			{k.x, k.y, k.z - 1},
			{k.x, k.y, k.z + 1},
		}
		for _, p := range prospects {
			if exterior[p] {
				exteriorSurface++
			}
		}
	}

	return uint(exteriorSurface)
}

type D3 struct {
	x, y, z int
}
