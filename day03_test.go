package adventofcode2022

import "testing"

func TestDay03Part1Example(t *testing.T) {
	const want = 157
	lines := linesFromFilename(t, exampleFilename(3))
	got := Day03(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part1(t *testing.T) {
	const want = 8109
	lines := linesFromFilename(t, filename(3))
	got := Day03(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// Benchmark is identical for part 1 and part 2 (same algorithm)
func BenchmarkDay03Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(3))
	for range b.N {
		_ = Day03(lines, true)
	}
}

func TestDay03Part2Example(t *testing.T) {
	const want = 70
	lines := linesFromFilename(t, exampleFilename(3))
	got := Day03(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part2(t *testing.T) {
	const want = 2738
	lines := linesFromFilename(t, filename(3))
	got := Day03(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(3))
	for range b.N {
		_ = Day03(lines, false)
	}
}
