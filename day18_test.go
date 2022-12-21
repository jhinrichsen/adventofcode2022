package adventofcode2022

import "testing"

func TestDay18Part1Example(t *testing.T) {
	const want = 64
	lines, err := linesFromFilename(exampleFilename(18))
	if err != nil {
		t.Fatal(err)
	}
	got := Day18(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want = 3550
	lines, err := linesFromFilename(filename(18))
	if err != nil {
		t.Fatal(err)
	}
	got := Day18(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2Example(t *testing.T) {
	const want = 58
	lines, err := linesFromFilename(exampleFilename(18))
	if err != nil {
		t.Fatal(err)
	}
	got := Day18(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	const want = 3304 // TODO too high
	lines, err := linesFromFilename(filename(18))
	if err != nil {
		t.Fatal(err)
	}
	got := Day18(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
