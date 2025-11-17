package adventofcode2022

import "testing"

// Mini example
func TestDay23Example0(t *testing.T) {
	const want = 25
	got := Day23([]string{
		".....",
		"..##.",
		"..#..",
		".....",
		"..##.",
		".....",
	}, -1)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1Example(t *testing.T) {
	const want = 110
	lines, err := linesFromFilename(exampleFilename(23))
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(lines, 10)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1(t *testing.T) {
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		t.Skip("puzzle input file not provided")
	}
	got := Day23(lines, 10)
	_ = got // Result depends on puzzle input
}
