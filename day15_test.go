package adventofcode2022

import "testing"

func TestDay15Part1Example(t *testing.T) {
	const want uint = 26
	lines, err := linesFromFilename(exampleFilename(15))
	if err != nil {
		t.Fatal(err)
	}
	got := Day15(lines, 10, true) // Example uses y=10
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay15Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(15))
	if err != nil {
		b.Skip("input file not found")
	}
	b.ResetTimer()
	for range b.N {
		_ = Day15(lines, 2000000, true) // Real puzzle uses y=2000000
	}
}
