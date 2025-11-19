package adventofcode2022

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func linesFromFilename(tb testing.TB, filename string) []string {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	lines := linesFromReader(tb, f)
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}
	return lines
}

func linesFromReader(tb testing.TB, r io.Reader) []string {
	tb.Helper()
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		tb.Fatal(err)
	}
	return lines
}

func numbersFromFilename(tb testing.TB, filename string) []int {
	tb.Helper()
	lines := linesFromFilename(tb, filename)
	numbers := make([]int, 0, len(lines))
	for _, line := range lines {
		var n int
		_, err := fmt.Sscanf(line, "%d", &n)
		if err != nil {
			continue // Skip lines that don't parse as integers
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func exampleFilename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", int(day))
}

func filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d.txt", int(day))
}

// Tests for helper functions

type CompareError struct {
	line int
	want string
	got  string
}

func (a CompareError) Error() string {
	return fmt.Sprintf("error comparing line %d: want %q but got %q",
		a.line, a.want, a.got)
}

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

func TestCompareEqual(t *testing.T) {
	err := compare(
		strings.NewReader(strings.Join([]string{
			"line 0",
			"line 1",
			"line 2",
		}, "\n")),
		strings.NewReader(strings.Join([]string{
			"line 0",
			"line 1",
			"line 2",
		}, "\n")),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCompareDifferent(t *testing.T) {
	err := compare(
		strings.NewReader(strings.Join([]string{
			"line 0",
			"line 1",
			"line 2",
		}, "\n")),
		strings.NewReader(strings.Join([]string{
			"line 0",
			"line 1",
			"Aller guten Dinge sind drei",
		}, "\n")))
	if err == nil {
		t.Fatalf("want error but got nil")
	}
	var ce *CompareError
	if errors.As(err, &ce) {
		if ce.line != 3 {
			t.Fatalf("want %d but got %d", 3, ce.line)
		}
	} else {
		t.Fatalf("cannot As into CompareError: %+v", err)
	}
}
