package adventofcode2022

import "testing"

func TestDay16Part1Example(t *testing.T) {
	const want uint = 1651
	lines, err := linesFromFilename(exampleFilename(16))
	if err != nil {
		t.Fatal(err)
	}
	got := Day16(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay16Part1(t *testing.T) {
	const want uint = 2080
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		t.Fatal(err)
	}
	got := Day16(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay16Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(16))
	if err != nil {
		b.Skip("input file not found")
	}
	b.ResetTimer()
	for range b.N {
		_ = Day16(lines, true)
	}
}
