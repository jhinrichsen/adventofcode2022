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

/*
TODO
func TestDay23Example(t *testing.T) {
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
*/
