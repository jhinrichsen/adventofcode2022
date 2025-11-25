package adventofcode2022

import "testing"

func TestDay07ExamplePart1(t *testing.T) {
	testLines(t, 7, exampleFilename, true, Day07, 95437)
}

func TestDay07Part1(t *testing.T) {
	testLines(t, 7, filename, true, Day07, 1428881)
}

func TestDay07ExamplePart2(t *testing.T) {
	testLines(t, 7, exampleFilename, false, Day07, 24933642)
}

func TestDay07Part2(t *testing.T) {
	testLines(t, 7, filename, false, Day07, 10475598)
}

func BenchmarkDay07Part1(b *testing.B) {
	benchLines(b, 7, true, Day07)
}

func BenchmarkDay07Part2(b *testing.B) {
	benchLines(b, 7, false, Day07)
}
