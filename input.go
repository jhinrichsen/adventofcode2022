package adventofcode2022

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func linesFromFilename(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func exampleFilename(day int) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func filename(day int) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}

// linesAsNumber converts strings into integer.
func linesAsNumbers(lines []string) ([]int, error) {
	var is []int
	for i := range lines {
		n, err := strconv.Atoi(lines[i])
		if err != nil {
			msg := "error in line %d: cannot convert %q to number"
			return is, fmt.Errorf(msg, i, lines[i])
		}
		is = append(is, n)
	}
	return is, nil
}

func numbersFromFilename(filename string) ([]int, error) {
	ls, err := linesFromFilename(filename)
	if err != nil {
		return nil, err
	}
	return linesAsNumbers(ls)
}

// ParseCommaSeparatedNumbers returns a partial list in case parsing fails.
func ParseCommaSeparatedNumbers(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	is := make([]int, len(parts))
	var err error
	for i := range parts {
		is[i], err = strconv.Atoi(parts[i])
		if err != nil {
			return is, err
		}
	}
	return is, nil
}

type CompareError struct {
	// the first line is 1, not 0. At least in vi. Can't go wrong following
	// vi.
	line int
	want string
	got  string
}

func (a CompareError) Error() string {
	return fmt.Sprintf("error comparing line %d: want %q but got %q",
		a.line, a.want, a.got)
}

// compare returns true if two reader deliver the same (line based) content.
func compare(want, got io.Reader) error {
	wsc := bufio.NewScanner(want)
	gsc := bufio.NewScanner(got)

	for line := 1; ; line++ {
		wb := wsc.Scan()
		gb := gsc.Scan()
		if wb != gb {
			return &CompareError{line, "", ""}
		}
		if !(wb || gb) {
			break
		}
		wt := wsc.Text()
		gt := gsc.Text()
		if wt != gt {
			return &CompareError{line, wt, gt}
		}
	}
	return nil
}
