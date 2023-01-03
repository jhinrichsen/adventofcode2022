package adventofcode2022

import "testing"

func TestDay08Part1Example(t *testing.T) {
	const want = 21
	lines, err := linesFromFilename(exampleFilename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part1(t *testing.T) {
	const want = 1835
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2Example(t *testing.T) {
	const want = 8
	lines, err := linesFromFilename(exampleFilename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2(t *testing.T) {
	const want = 263670
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
