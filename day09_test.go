package adventofcode2022

import "testing"

func TestDay09Example(t *testing.T) {
	const want = 13
	lines, err := linesFromFilename(exampleFilename(9))
	if err != nil {
		t.Fatal(err)
	}
	got := Day09(lines)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay09(t *testing.T) {
	const want = 6212
	lines, err := linesFromFilename(filename(9))
	if err != nil {
		t.Fatal(err)
	}
	got := Day09(lines)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
