package adventofcode2022

import (
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

func Day12(lines []string, part1 bool) int {
	var (
		dimX       = len(lines[0])
		dimY       = len(lines)
		world      = simple.NewDirectedGraph()
		start, end graph.Node

		idx = func(x, y int) int64 { // 2d -> 1d
			return int64(y*dimX + x)
		}

		xy = func(idx int64) (int, int) { // 1d -> 2d
			return int(idx) / dimX, int(idx) % dimX
		}

		mk = func(x, y int) graph.Node {
			id := idx(x, y)
			f := simple.Node(id)
			return f
		}

		connect = func(x0, y0, x1, y1 int) {
			f := mk(x0, y0)
			t := mk(x1, y1)
			world.SetEdge(simple.Edge{F: f, T: t})
		}

		nirvana = func(x, y int) bool {
			return x < 0 || x >= dimX || y < 0 || y >= dimY
		}

		height2D = func(x, y int) int { // 2d -> h
			return int(lines[y][x])
		}

		height1D = func(idx int64) int { // 1d -> h
			return height2D(xy(idx))
		}

		prospect = func(x0, y0, dx, dy int) {
			x1 := x0 + dx
			y1 := y0 + dy
			if nirvana(x1, y1) {
				return
			}
			h0 := height2D(x0, y0)
			h1 := height2D(x1, y1)

			// step down unlimited, step up one max
			if h1-h0 <= 1 {
				connect(x0, y0, x1, y1)
			}
		}
	)

	// phase 1: locate start and end position
	// y₀? nope, 'invalid character U+2080 '₀'
	for y := range lines {
		for x := range lines[y] {
			cell := lines[y][x]
			if cell == 'S' {
				start = simple.Node(idx(x, y))
				lines[y] = strings.Replace(lines[y], "S", "a", 1)
			} else if cell == 'E' {
				end = simple.Node(idx(x, y))
				lines[y] = strings.Replace(lines[y], "E", "z", 1)
			}

			if start != nil && end != nil {
				goto short_circuit
			}
		}
	}
short_circuit:
	if start == nil || end == nil {
		panic("cannot find required S or E")
	}

	// phase 2: build edges
	for y0 := range lines {
		for x0 := range lines[y0] {
			prospect(x0, y0, 0, -1) // north
			prospect(x0, y0, 0, 1)  // south
			prospect(x0, y0, 1, 0)  // east
			prospect(x0, y0, -1, 0) // west
		}
	}

	pth, _ := path.DijkstraFrom(start, world).To(end.ID())

	if part1 {
		return len(pth) - 1 // we need steps which is one less than nodes
	}

	// count up to the last step that has the same height as the first step
	i := 1
	for ; i < len(pth); i++ {
		h0 := height1D(pth[i-1].ID())
		h1 := height1D(pth[i].ID())
		if h1 != h0 {
			break
		}
	}
	return len(pth) - i - 1
}
