package adventofcode2022

import "testing"

func TestDay13Part1Example(t *testing.T) {
	testLines(t, 13, exampleFilename, true, Day13, uint(13))
}

func TestDay13Part1(t *testing.T) {
	testLines(t, 13, filename, true, Day13, uint(5808))
}

func TestDay13Part2Example(t *testing.T) {
	testLines(t, 13, exampleFilename, false, Day13, uint(140))
}

func TestDay13Part2(t *testing.T) {
	testLines(t, 13, filename, false, Day13, uint(22713))
}

func BenchmarkDay13Part1(b *testing.B) {
	benchLines(b, 13, true, Day13)
}

func BenchmarkDay13Part2(b *testing.B) {
	benchLines(b, 13, false, Day13)
}
