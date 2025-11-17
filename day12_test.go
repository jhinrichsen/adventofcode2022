package adventofcode2022

import "testing"

func TestDay12Part1Example(t *testing.T) {
	const want uint = 31
	lines, err := linesFromFilename(exampleFilename(12))
	if err != nil {
		t.Fatal(err)
	}
	got := Day12(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part1(t *testing.T) {
	const want uint = 528
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		t.Fatal(err)
	}
	got := Day12(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part2Example(t *testing.T) {
	const want uint = 29
	lines, err := linesFromFilename(exampleFilename(12))
	if err != nil {
		t.Fatal(err)
	}
	got := Day12(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part2(t *testing.T) {
	const want uint = 522
	lines, err := linesFromFilename(filename(12))
	if err != nil {
		t.Fatal(err)
	}
	got := Day12(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
