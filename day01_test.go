package adventofcode2022

import "testing"

func TestDay01Part1Example(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(1))
	got := Day01(lines, 1)
	const want = 24000
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01(t *testing.T) {
	lines := linesFromFilename(t, filename(1))
	got := Day01(lines, 1)
	const want = 69177
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part2Example(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(1))
	got := Day01(lines, 3)
	const want = 45000
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay01Part2(t *testing.T) {
	lines := linesFromFilename(t, filename(1))
	got := Day01(lines, 3)
	const want = 207456
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(1))
	for b.Loop() {
		_ = Day01(lines, 1)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(1))
	for b.Loop() {
		_ = Day01(lines, 3)
	}
}
