package adventofcode2022

import (
	"fmt"
	"testing"
)

func example1Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example1.txt", int(day))
}

func example2Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example2.txt", int(day))
}

func TestDay09Part1Example(t *testing.T) {
	testLines(t, 9, example1Filename, true, Day09, 13)
}

func TestDay09Part1(t *testing.T) {
	testLines(t, 9, filename, true, Day09, 6044)
}

func BenchmarkDay09Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(9))
	for b.Loop() {
		_ = Day09(lines, true)
	}
}

func TestDay09Part2Example1(t *testing.T) {
	testLines(t, 9, example1Filename, false, Day09, 1)
}

func TestDay09Part2Example2(t *testing.T) {
	testLines(t, 9, example2Filename, false, Day09, 36)
}

func TestDay09Part2(t *testing.T) {
	testLines(t, 9, filename, false, Day09, 2384)
}

func BenchmarkDay09Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(9))
	for b.Loop() {
		_ = Day09(lines, false)
	}
}
