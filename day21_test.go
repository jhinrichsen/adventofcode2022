package adventofcode2022

import (
	"testing"
)

func TestDay21Part1Example(t *testing.T) {
	const want = 152
	lines := linesFromFilename(t, exampleFilename(21))
	got := Day21(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part1(t *testing.T) {
	const want = 194501589693264
	lines := linesFromFilename(t, filename(21))
	got := Day21(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay21Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(21))
	for range b.N {
		Day21(lines, true)
	}
}

func TestDay21Part2Example(t *testing.T) {
	const want = 301
	lines := linesFromFilename(t, exampleFilename(21))
	got := Day21(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part2(t *testing.T) {
	const want = 3887609741189
	lines := linesFromFilename(t, filename(21))
	got := Day21(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay21Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(21))
	for range b.N {
		Day21(lines, false)
	}
}
