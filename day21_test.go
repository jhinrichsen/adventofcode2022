package adventofcode2022

import (
	"testing"
)

func TestDay21Part1Example(t *testing.T) {
	testLines(t, 21, exampleFilename, true, Day21, 152)
}

func TestDay21Part1(t *testing.T) {
	testLines(t, 21, filename, true, Day21, 194501589693264)
}

func TestDay21Part2Example(t *testing.T) {
	testLines(t, 21, exampleFilename, false, Day21, 301)
}

func TestDay21Part2(t *testing.T) {
	testLines(t, 21, filename, false, Day21, 3887609741189)
}

func BenchmarkDay21Part1(b *testing.B) {
	benchLines(b, 21, true, Day21)
}

func BenchmarkDay21Part2(b *testing.B) {
	benchLines(b, 21, false, Day21)
}
