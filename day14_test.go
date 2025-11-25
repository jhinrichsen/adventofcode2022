package adventofcode2022

import "testing"

func TestDay14Part1Example(t *testing.T) {
	testLines(t, 14, exampleFilename, true, Day14, uint(24))
}

func TestDay14Part1(t *testing.T) {
	testLines(t, 14, filename, true, Day14, uint(779))
}

func TestDay14Part2Example(t *testing.T) {
	testLines(t, 14, exampleFilename, false, Day14, uint(93))
}

func TestDay14Part2(t *testing.T) {
	testLines(t, 14, filename, false, Day14, uint(27426))
}

func BenchmarkDay14Part1(b *testing.B) {
	benchLines(b, 14, true, Day14)
}

func BenchmarkDay14Part2(b *testing.B) {
	benchLines(b, 14, false, Day14)
}
