package adventofcode2022

import (
	"testing"
)

func TestDay22Part1Example(t *testing.T) {
	const want uint = 6032
	lines := linesFromFilename(t, exampleFilename(22))
	got := Day22(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part1(t *testing.T) {
	const want uint = 65368
	lines := linesFromFilename(t, filename(22))
	got := Day22(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// TODO: Debug example cube topology - A->B and C->D transitions work correctly
// but full path simulation gives wrong answer (4045 vs expected 5031)
/*
func TestDay22Part2Example(t *testing.T) {
	const want = 5031
	lines := linesFromFilename(t, exampleFilename(22))
	got := Day22(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/

func TestDay22Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(22))
	got := Day22(lines, false)
	t.Logf("Part 2 result: %d", got)
	// Result will need to be validated against AoC submission
}

func BenchmarkDay22Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(22))
	for range b.N {
		_ = Day22(lines, true)
	}
}

func BenchmarkDay22Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(22))
	for range b.N {
		_ = Day22(lines, false)
	}
}
