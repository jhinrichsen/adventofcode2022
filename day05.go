package adventofcode2022

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Day05(input []byte, part1 bool) (string, error) {
	const width = len("[A] ")

	lines := bytes.Split(input, []byte{'\n'})
	if len(lines) == 0 {
		return "", fmt.Errorf("empty input")
	}

	lineIdx := 0
	line := string(lines[lineIdx])
	lineIdx++

	n := (len(line) + 1) / len("[A] ") // number of stacks
	crate := func(line string, n int) byte {
		return line[n*width+1]
	}
	st := make([][]byte, n)

	crateLine := func(line string) {
		for i := 0; i < n; i++ {
			// skip empty crates/ space character
			c := crate(line, i)
			if c != ' ' {
				st[i] = append(st[i], c)
			}
		}
	}

	crateLine(line)
	for lineIdx < len(lines) {
		line = string(lines[lineIdx])
		lineIdx++
		if len(line) > 1 && line[1] == '1' { // reached end of stacks
			// skip empty line
			lineIdx++
			break
		}
		if len(line) > 0 {
			crateLine(line)
		}
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

	for lineIdx < len(lines) {
		line = string(lines[lineIdx])
		lineIdx++
		if len(line) == 0 {
			continue
		}

		cmds := strings.Fields(line)
		if len(cmds) < 6 {
			continue
		}

		n, err := strconv.Atoi(cmds[1])
		if err != nil {
			return "", fmt.Errorf("error parsing move %q", cmds[1])
		}
		from, err := strconv.Atoi(cmds[3])
		if err != nil {
			return "", fmt.Errorf("error parsing from %q", cmds[3])
		}
		into, err := strconv.Atoi(cmds[5])
		if err != nil {
			return "", fmt.Errorf("error parsing to %q", cmds[5])
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
		if len(st[i]) > 0 {
			sb.WriteByte(st[i][len(st[i])-1])
		}
	}
	return sb.String(), nil
}
