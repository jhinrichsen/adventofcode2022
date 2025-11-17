package adventofcode2022

import "testing"

func TestDay24Part1Example(t *testing.T) {
	const want uint = 18
	lines, err := linesFromFilename(exampleFilename(24))
	if err != nil {
		t.Fatal(err)
	}
	got := Day24(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part1(t *testing.T) {
	lines, err := linesFromFilename(filename(24))
	if err != nil {
		t.Skip("puzzle input file not provided")
	}
	got := Day24(lines, true)
	_ = got // Result depends on puzzle input
}
