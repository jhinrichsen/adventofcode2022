package adventofcode2022

import (
	"path/filepath"
	"strconv"
	"strings"
)

func Day07(lines []string) int {
	const limit = 100000

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

	// filter and sum
	sum := 0
	for _, size := range ds {
		if size <= limit {
			sum += size
		}
	}
	return sum
}
