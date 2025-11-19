package adventofcode2022

import "testing"

func TestDay13Part1Example(t *testing.T) {
	const want uint = 13
	lines := linesFromFilename(t, exampleFilename(13))
	got := Day13(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part1(t *testing.T) {
	const want uint = 5808
	lines := linesFromFilename(t, filename(13))
	got := Day13(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part2Example(t *testing.T) {
	const want uint = 140
	lines := linesFromFilename(t, exampleFilename(13))
	got := Day13(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part2(t *testing.T) {
	const want uint = 22713
	lines := linesFromFilename(t, filename(13))
	got := Day13(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(13))
	for range b.N {
		_ = Day13(lines, true)
	}
}

func BenchmarkDay13Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(13))
	for range b.N {
		_ = Day13(lines, false)
	}
}
