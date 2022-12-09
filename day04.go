package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

func Day04(lines []string) (int, error) {
	parse := func(line string) (int, int, int, int, error) {
		var err error
		l0 := strings.Split(line, ",")
		l1 := strings.Split(l0[0], "-")
		l2 := strings.Split(l0[1], "-")
		ss := []string{l1[0], l1[1], l2[0], l2[1]}
		var ns [4]int
		for i := range ns {
			ns[i], err = strconv.Atoi(ss[i])
			if err != nil {
				msg := "error parsing number %q: %w"
				err = fmt.Errorf(msg, ss[i], err)
				break
			}
		}
		return ns[0], ns[1], ns[2], ns[3], err
	}

	count := 0
	for i, line := range lines {
		x1, x2, x3, x4, err := parse(line)
		if err != nil {
			return 0, fmt.Errorf("error parsing line %d: %w",
				i+1, err)
		}
		if Contains(x1, x2, x3, x4) {
			count++
		}
	}
	return count, nil
}

// Contains returns true if the range a [a1..a2] is within b [b1..b2] or vice
// versa.
func Contains(a1, a2, b1, b2 int) bool {
	return a1 >= b1 && a2 <= b2 || // a within b
		a1 <= b1 && a2 >= b2 // b within a
}
