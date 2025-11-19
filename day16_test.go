package adventofcode2022

import "testing"

func TestDay16Part1Example(t *testing.T) {
	const want uint = 1651
	lines := linesFromFilename(t, exampleFilename(16))
	got := Day16(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay16Part1(t *testing.T) {
	const want uint = 2080
	lines := linesFromFilename(t, filename(16))
	got := Day16(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay16Part2Example(t *testing.T) {
	const want uint = 1707
	lines := linesFromFilename(t, exampleFilename(16))
	got := Day16(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay16Part2(t *testing.T) {
	const want uint = 2752
	lines := linesFromFilename(t, filename(16))
	got := Day16(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay16Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(16))
	for range b.N {
		_ = Day16(lines, true)
	}
}

func BenchmarkDay16Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(16))
	for range b.N {
		_ = Day16(lines, false)
	}
}
