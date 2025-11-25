package adventofcode2022

import "testing"

func TestDay01Part1Example(t *testing.T) {
	testLines(t, 1, exampleFilename, true, Day01, uint(24000))
}

func TestDay01Part1(t *testing.T) {
	testLines(t, 1, filename, true, Day01, uint(69177))
}

func TestDay01Part2Example(t *testing.T) {
	testLines(t, 1, exampleFilename, false, Day01, uint(45000))
}

func TestDay01Part2(t *testing.T) {
	testLines(t, 1, filename, false, Day01, uint(207456))
}

func BenchmarkDay01Part1(b *testing.B) {
	benchLines(b, 1, true, Day01)
}

func BenchmarkDay01Part2(b *testing.B) {
	benchLines(b, 1, false, Day01)
}
