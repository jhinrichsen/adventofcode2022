package adventofcode2022

import "testing"

func TestDay12Part1Example(t *testing.T) {
	testLines(t, 12, exampleFilename, true, Day12, uint(31))
}

func TestDay12Part1(t *testing.T) {
	testLines(t, 12, filename, true, Day12, uint(528))
}

func TestDay12Part2Example(t *testing.T) {
	testLines(t, 12, exampleFilename, false, Day12, uint(29))
}

func TestDay12Part2(t *testing.T) {
	testLines(t, 12, filename, false, Day12, uint(522))
}

func BenchmarkDay12Part1(b *testing.B) {
	benchLines(b, 12, true, Day12)
}

func BenchmarkDay12Part2(b *testing.B) {
	benchLines(b, 12, false, Day12)
}
