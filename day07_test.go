package adventofcode2022

import "testing"

func TestDay07ExamplePart1(t *testing.T) {
	const want = 95437
	lines := linesFromFilename(t, exampleFilename(7))
	got := Day07(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = 1428881
	lines := linesFromFilename(t, filename(7))
	got := Day07(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07ExamplePart2(t *testing.T) {
	const want = 24933642
	lines := linesFromFilename(t, exampleFilename(7))
	got := Day07(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	const want = 10475598
	lines := linesFromFilename(t, filename(7))
	got := Day07(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(7))
	for range b.N {
		_ = Day07(lines, true)
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(7))
	for range b.N {
		_ = Day07(lines, false)
	}
}
