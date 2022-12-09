package adventofcode2022

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func parseStacks(r io.Reader) []string {
	var ss []string
	return ss
}

func Day05(r io.Reader) (string, error) {
	const width = len("[A] ")

	sc := bufio.NewScanner(r)
	// read the first line and use as template for setup
	ok, line := sc.Scan(), sc.Text()
	if !ok {
		return "", fmt.Errorf("error reading from reader")
	}

	n := (len(line) + 1) / len("[A] ") // number of stacks
	crate := func(n int) byte {
		return line[n*width+1]
	}
	st := make([][]byte, n)
	// macro like function
	crateLine := func() {
		for i := 0; i < n; i++ {
			// skip empty crates/ space character
			c := crate(i)
			if c != ' ' {
				st[i] = append(st[i], c)
			}
		}
	}

	crateLine()
	for sc.Scan() {
		line = sc.Text()
		if line[1] == '1' { // reached end of stacks, switch group
			// read next, empty line
			sc.Scan()
			break
		}
		crateLine()
	}

	for i := range st {
		Reverse(st[i])
	}

	move := func(n, from, into int) {
		for i := 0; i < n; i++ {
			// this section creates no garbage
			// pop from
			last := len(st[from]) - 1
			n := st[from][last]
			st[from] = st[from][:last]

			// push into
			st[into] = append(st[into], n)
		}
	}
	for sc.Scan() {
		cmds := strings.Fields(sc.Text())
		n, err := strconv.Atoi(cmds[1])
		if err != nil {
			return "", fmt.Errorf("error parsing move %q",
				cmds[1])
		}
		from, err := strconv.Atoi(cmds[3])
		if err != nil {
			return "", fmt.Errorf("error parsing from %q",
				cmds[3])
		}
		into, err := strconv.Atoi(cmds[5])
		if err != nil {
			return "", fmt.Errorf("error parsing to %q",
				cmds[5])
		}
		move(n, from-1, into-1) // stacks are 1-based
	}

	// gather last of each stack
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(st[i][len(st[i])-1])
	}
	return sb.String(), nil
}
