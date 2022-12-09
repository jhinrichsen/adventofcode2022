package adventofcode2022

import (
	"fmt"
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
				fmt.Printf("wd cd %s\n", ss[2])
				if ss[2][0] == '/' {
					wd = ss[2]
				} else {
					wd = filepath.Join(wd, ss[2])
					fmt.Printf("wd before clean %s\n", wd)
					wd = filepath.Clean(wd)
					fmt.Printf("wd after clean %s\n", wd)
				}
				fmt.Printf("wd switching to %s\n", wd)
			}
		case "dir": // directory
			// record directories so that intermediate directories
			// without files are accounted for
			dirs[filepath.Join(wd, ss[1])] = true
		default: // filesize
			n, err := strconv.Atoi(ss[0])
			if err != nil {
				panic("bad")
			}
			f := filepath.Join(wd, ss[1])
			fmt.Printf("wd registering %s\n", f)
			sizes[f] = n
			dirs[wd] = true
			fmt.Printf("file %s: %d\n", f, n)
		}
	}

	/*
		for k := range dirs {
			fmt.Printf("unique %s: \n", k)
		}
	*/

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
		// fmt.Printf("%s: %d\n", dir, size)
		if size <= limit {
			sum += size
		}
	}
	return sum
}
