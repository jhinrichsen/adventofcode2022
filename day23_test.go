package adventofcode2022

import "testing"

// Mini example - Part 2
func TestDay23Part2Example0(t *testing.T) {
	const want uint = 4
	got := Day23([]string{
		".....",
		"..##.",
		"..#..",
		".....",
		"..##.",
		".....",
	}, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1Example(t *testing.T) {
	testLines(t, 23, exampleFilename, true, Day23, uint(110))
}

func TestDay23Part1(t *testing.T) {
	testLines(t, 23, filename, true, Day23, uint(3689))
}

func TestDay23Part2Example(t *testing.T) {
	testLines(t, 23, exampleFilename, false, Day23, uint(20))
}

func TestDay23Part2(t *testing.T) {
	testLines(t, 23, filename, false, Day23, uint(965))
}

func BenchmarkDay23Part1(b *testing.B) {
	benchLines(b, 23, true, Day23)
}

func BenchmarkDay23Part2(b *testing.B) {
	benchLines(b, 23, false, Day23)
}
