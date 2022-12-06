package adventofcode2021

import "testing"

func TestDay03Part1Example(t *testing.T) {
	const want = 157
	lines, err := linesFromFilename(exampleFilename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Part1(t *testing.T) {
	const want = 7746
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// Benchmark is identical for part 1 and part 2 (same algorithm)
func BenchmarkDay03(b *testing.B) {
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day03(lines)
	}
}

/*
func TestDay03Part2Example(t *testing.T) {
	const want = 12
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
	const want = 14979
	lines, err := linesFromFilename(filename(3))
	if err != nil {
		t.Fatal(err)
	}
	got := Day03(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/
