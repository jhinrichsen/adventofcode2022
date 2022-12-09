package adventofcode2022

import (
	"math"
	"path/filepath"
	"strconv"
	"strings"
)

func Day07(lines []string, part1 bool) int {
	var wd string                 // working directory
	dirs := make(map[string]bool) // list of directories
	sizes := make(map[string]int) // filename -> size

	for _, line := range lines {
		ss := strings.Fields(line)
		switch ss[0] {
		case "$": // command
			if ss[1] == "cd" {
				if ss[2][0] == '/' {
					wd = ss[2]
				} else {
					wd = filepath.Join(wd, ss[2])
					wd = filepath.Clean(wd)
				}
			}
		case "dir": // directory
			// record directories so that intermediate directories
			// without files are accounted for
			dirs[filepath.Join(wd, ss[1])] = true
		default: // filesize
			n, _ := strconv.Atoi(ss[0])
			f := filepath.Join(wd, ss[1])
			sizes[f] = n
			dirs[wd] = true
		}
	}

	// calculate directory sizes
	ds := make(map[string]int, len(dirs))
	for dir := range dirs {
		for filename, size := range sizes {
			if strings.HasPrefix(filepath.Dir(filename), dir) {
				ds[dir] += size
			}
		}
	}

	if part1 {
		const limit = 100000

		// filter and sum
		sum := 0
		for _, size := range ds {
			if size <= limit {
				sum += size
			}
		}
		return sum
	}

	// part 2
	const (
		available = 70_000_000
		required  = 30_000_000
	)
	used := ds["/"]
	free := available - used
	scratch := required - free
	if scratch < 0 {
		return 0
	}
	var n int
	Δmin := math.MaxInt
	for _, v := range ds {
		Δ := v - scratch
		if Δ > 0 && Δ < Δmin {
			Δmin = Δ
			n = v
		}
	}
	return n
}
