package adventofcode2022

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Day05(r io.Reader, part1 bool) (string, error) {
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

	move1 := func(n, from, into int) {
		for i := 0; i < n; i++ {
			// this section creates no garbage
			// pop from
			last := len(st[from]) - 1
			pick := st[from][last]
			st[from] = st[from][:last]

			// push into
			st[into] = append(st[into], pick)
		}
	}

	moveN := func(n, from, into int) {
		last := len(st[from])
		pick := st[from][last-n:]
		st[from] = st[from][:last-n]
		st[into] = append(st[into], pick...)
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

		if part1 {
			move1(n, from-1, into-1) // stacks are 1-based
		} else {
			moveN(n, from-1, into-1) // stacks are 1-based
		}
	}

	// gather last of each stack
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(st[i][len(st[i])-1])
	}
	return sb.String(), nil
}
