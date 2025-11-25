package adventofcode2022

import "testing"

func TestDay15Part1Example(t *testing.T) {
	const want uint = 26
	lines := linesFromFilename(t, exampleFilename(15))
	got := Day15(lines, 10, 20, true) // Example uses y=10, max=20
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay15Part1(t *testing.T) {
	const want uint = 5607466
	lines := linesFromFilename(t, filename(15))
	got := Day15(lines, 2000000, 4000000, true) // Real puzzle uses y=2000000, max=4000000
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay15Part2Example(t *testing.T) {
	const want uint = 56000011
	lines := linesFromFilename(t, exampleFilename(15))
	got := Day15(lines, 10, 20, false) // Example uses max=20
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay15Part2(t *testing.T) {
	const want uint = 12543202766584
	lines := linesFromFilename(t, filename(15))
	got := Day15(lines, 2000000, 4000000, false) // Real puzzle uses max=4000000
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay15Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(15))
	for b.Loop() {
		_ = Day15(lines, 2000000, 4000000, true) // Real puzzle uses y=2000000, max=4000000
	}
}

func BenchmarkDay15Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(15))
	for b.Loop() {
		_ = Day15(lines, 2000000, 4000000, false) // Real puzzle uses max=4000000
	}
}
