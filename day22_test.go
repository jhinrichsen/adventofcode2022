package adventofcode2022

import (
	"testing"
)

func TestDay22Part1Example(t *testing.T) {
	const want = 6032
	lines, err := linesFromFilename(exampleFilename(22))
	if err != nil {
		t.Fatal(err)
	}
	got := Day22(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

/*
func TestDay22Part1(t *testing.T) {
	const want = 194501589693264
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		t.Fatal(err)
	}
	got := Day22(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay22Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day22(lines, true)
	}
}

func TestDay22Part2Example(t *testing.T) {
	const want = 58
	lines, err := linesFromFilename(exampleFilename(22))
	if err != nil {
		t.Fatal(err)
	}
	got := Day22(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay22Part2(t *testing.T) {
	const want = 3304 // too high
	lines, err := linesFromFilename(filename(22))
	if err != nil {
		t.Fatal(err)
	}
	got := Day22(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
*/
