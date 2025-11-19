package adventofcode2022

import "testing"

func TestDay14Part1Example(t *testing.T) {
	const want uint = 24
	lines := linesFromFilename(t, exampleFilename(14))
	got := Day14(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part1(t *testing.T) {
	const want uint = 779
	lines := linesFromFilename(t, filename(14))
	got := Day14(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part2Example(t *testing.T) {
	const want uint = 93
	lines := linesFromFilename(t, exampleFilename(14))
	got := Day14(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part2(t *testing.T) {
	const want uint = 27426
	lines := linesFromFilename(t, filename(14))
	got := Day14(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay14Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(14))
	for range b.N {
		_ = Day14(lines, true)
	}
}

func BenchmarkDay14Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(14))
	for range b.N {
		_ = Day14(lines, false)
	}
}
