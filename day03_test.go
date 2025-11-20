package adventofcode2022

import "testing"

func TestDay03Part1Example(t *testing.T) {
	testLines(t, 3, exampleFilename, true, Day03, 157)
}

func TestDay03Part1(t *testing.T) {
	testLines(t, 3, filename, true, Day03, 8109)
}

func TestDay03Part2Example(t *testing.T) {
	testLines(t, 3, exampleFilename, false, Day03, 70)
}

func TestDay03Part2(t *testing.T) {
	testLines(t, 3, filename, false, Day03, 2738)
}

func BenchmarkDay03Part1(b *testing.B) {
	benchLines(b, 3, true, Day03)
}

func BenchmarkDay03Part2(b *testing.B) {
	benchLines(b, 3, false, Day03)
}
