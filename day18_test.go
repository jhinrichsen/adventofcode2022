package adventofcode2022

import "testing"

func TestDay18Part1Example(t *testing.T) {
	testLines(t, 18, exampleFilename, true, Day18, uint(64))
}

func TestDay18Part1(t *testing.T) {
	testLines(t, 18, filename, true, Day18, uint(3550))
}

func TestDay18Part2Example(t *testing.T) {
	testLines(t, 18, exampleFilename, false, Day18, uint(58))
}

func TestDay18Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(18))
	got := Day18(lines, false)
	t.Logf("Day 18 Part 2: %d", got)
}

func BenchmarkDay18Part1(b *testing.B) {
	benchLines(b, 18, true, Day18)
}

func BenchmarkDay18Part2(b *testing.B) {
	benchLines(b, 18, false, Day18)
}
