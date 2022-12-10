package adventofcode2022

import "testing"

func TestDay08Example(t *testing.T) {
	const want = 21
	lines, err := linesFromFilename(exampleFilename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08(t *testing.T) {
	const want = 1779
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		t.Fatal(err)
	}
	got := Day08(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
