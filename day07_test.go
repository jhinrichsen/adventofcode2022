package adventofcode2022

import "testing"

func TestDay07Example(t *testing.T) {
	const want = 95437
	lines, err := linesFromFilename(exampleFilename(7))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07(t *testing.T) {
	const want = 1243729
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
