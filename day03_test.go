package adventofcode2022

import "testing"

func TestDay03Part1Example(t *testing.T) {
	const want = 157
	lines, err := linesFromFilename(exampleFilename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part1(t *testing.T) {
	const want = 8109
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// Benchmark is identical for part 1 and part 2 (same algorithm)
func BenchmarkDay03Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day03(lines, true)
	}
}

func TestDay03Part2Example(t *testing.T) {
	const want = 70
	lines, err := linesFromFilename(exampleFilename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part2(t *testing.T) {
	const want = 2738
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day03(lines, false)
	}
}
