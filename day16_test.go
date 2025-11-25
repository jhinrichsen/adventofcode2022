package adventofcode2022

import "testing"

func TestDay16Part1Example(t *testing.T) {
	testLines(t, 16, exampleFilename, true, Day16, uint(1651))
}

func TestDay16Part1(t *testing.T) {
	testLines(t, 16, filename, true, Day16, uint(2080))
}

func TestDay16Part2Example(t *testing.T) {
	testLines(t, 16, exampleFilename, false, Day16, uint(1707))
}

func TestDay16Part2(t *testing.T) {
	testLines(t, 16, filename, false, Day16, uint(2752))
}

func BenchmarkDay16Part1(b *testing.B) {
	benchLines(b, 16, true, Day16)
}

func BenchmarkDay16Part2(b *testing.B) {
	benchLines(b, 16, false, Day16)
}
