package adventofcode2022

import "testing"

func TestDay02Part1Example(t *testing.T) {
	testLines(t, 2, exampleFilename, true, Day02, 15)
}

func TestDay02Part1(t *testing.T) {
	testLines(t, 2, filename, true, Day02, 13052)
}

func TestDay02Part2Example(t *testing.T) {
	testLines(t, 2, exampleFilename, false, Day02, 12)
}

func TestDay02Part2(t *testing.T) {
	testLines(t, 2, filename, false, Day02, 13693)
}

func BenchmarkDay02Part1(b *testing.B) {
	benchLines(b, 2, true, Day02)
}

func BenchmarkDay02Part2(b *testing.B) {
	benchLines(b, 2, false, Day02)
}
