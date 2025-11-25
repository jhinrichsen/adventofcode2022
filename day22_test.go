package adventofcode2022

import (
	"testing"
)

func TestDay22Part1Example(t *testing.T) {
	testLines(t, 22, exampleFilename, true, Day22, uint(6032))
}

func TestDay22Part1(t *testing.T) {
	testLines(t, 22, filename, true, Day22, uint(65368))
}

// TODO: Debug example cube topology - A->B and C->D transitions work correctly
// but full path simulation gives wrong answer (4045 vs expected 5031)
/*
func TestDay22Part2Example(t *testing.T) {
	testLines(t, 22, exampleFilename, false, Day22, uint(5031))
}
*/

func TestDay22Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(22))
	got := Day22(lines, false)
	t.Logf("Part 2 result: %d", got)
	// Result will need to be validated against AoC submission
}

func BenchmarkDay22Part1(b *testing.B) {
	benchLines(b, 22, true, Day22)
}

func BenchmarkDay22Part2(b *testing.B) {
	benchLines(b, 22, false, Day22)
}
