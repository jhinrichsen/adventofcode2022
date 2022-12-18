package adventofcode2022

import (
	"fmt"
	"os"
)

func Day13(lines []string) int {
	var (
		out = os.Stdout
	)

	var pair int // 1-based pair index
	for i := 0; i < len(lines); i += 3 {
		pair++
		left := lines[i]
		right := lines[i+1]
		fmt.Fprintf(out, "== Pair %d ==\n", pair)
		fmt.Fprintf(out, "- Compare %v versus %v\n", left, right)
	}
	return 0
}
