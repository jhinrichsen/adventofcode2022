package adventofcode2022

import (
	"testing"
)

func TestDay21Part1Example(t *testing.T) {
	const want = 152
	lines, err := linesFromFilename(exampleFilename(21))
	if err != nil {
		t.Fatal(err)
	}
	got := Day21(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part1(t *testing.T) {
	const want = 194501589693264
	lines, err := linesFromFilename(filename(21))
	if err != nil {
		t.Fatal(err)
	}
	got := Day21(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay21Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(21))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day21(lines, true)
	}
}

/*
func TestDay21Part2Example(t *testing.T) {
	const want = 58
	lines, err := linesFromFilename(exampleFilename(21))
	if err != nil {
		t.Fatal(err)
	}
	got := Day21(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part2(t *testing.T) {
	const want = 3304 // too high
	lines, err := linesFromFilename(filename(21))
	if err != nil {
		t.Fatal(err)
	}
	got := Day21(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/
