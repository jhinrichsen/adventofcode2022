package adventofcode2022

import "testing"

func TestDay12Part1Example(t *testing.T) {
	const want uint = 31
	lines := linesFromFilename(t, exampleFilename(12))
	got := Day12(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part1(t *testing.T) {
	const want uint = 528
	lines := linesFromFilename(t, filename(12))
	got := Day12(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part2Example(t *testing.T) {
	const want uint = 29
	lines := linesFromFilename(t, exampleFilename(12))
	got := Day12(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part2(t *testing.T) {
	const want uint = 522
	lines := linesFromFilename(t, filename(12))
	got := Day12(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay12Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(12))
	for range b.N {
		_ = Day12(lines, true)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(12))
	for range b.N {
		_ = Day12(lines, false)
	}
}
