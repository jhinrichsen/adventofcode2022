package adventofcode2022

import "testing"

func TestDay18Part1Example(t *testing.T) {
	const want uint = 64
	lines := linesFromFilename(t, exampleFilename(18))
	got := Day18(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want uint = 3550
	lines := linesFromFilename(t, filename(18))
	got := Day18(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2Example(t *testing.T) {
	const want uint = 58
	lines := linesFromFilename(t, exampleFilename(18))
	got := Day18(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(18))
	got := Day18(lines, false)
	t.Logf("Day 18 Part 2: %d", got)
}

func BenchmarkDay18Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(18))
	for range b.N {
		_ = Day18(lines, true)
	}
}

func BenchmarkDay18Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(18))
	for range b.N {
		_ = Day18(lines, false)
	}
}
