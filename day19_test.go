package adventofcode2022

import "testing"

func TestDay19Part1Example(t *testing.T) {
	const want uint = 33
	lines, err := linesFromFilename(exampleFilename(19))
	if err != nil {
		t.Fatal(err)
	}
	got := Day19(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay19Part1(t *testing.T) {
	const want uint = 1356
	lines, err := linesFromFilename(filename(19))
	if err != nil {
		t.Fatal(err)
	}
	got := Day19(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
