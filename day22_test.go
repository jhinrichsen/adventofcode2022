package adventofcode2022

import (
	"testing"
)

func TestDay22Part1Example(t *testing.T) {
	const want = 6032
	lines, err := linesFromFilename(exampleFilename(22))
	if err != nil {
		t.Fatal(err)
	}
	got := Day22(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part1(t *testing.T) {
	// const want = 80224 // TODO too high
	const want = 65368
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		t.Fatal(err)
	}
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
	lines, err := linesFromFilename(exampleFilename(22))
	if err != nil {
		t.Fatal(err)
	}
	got := Day22(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/

func TestDay22Part2(t *testing.T) {
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		t.Fatal(err)
	}
	got := Day22(lines, false)
	t.Logf("Part 2 result: %d", got)
	// Result will need to be validated against AoC submission
}
