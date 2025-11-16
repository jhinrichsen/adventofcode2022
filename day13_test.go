package adventofcode2022

import "testing"

func TestDay13Part1Example(t *testing.T) {
	const want uint = 13
	lines, err := linesFromFilename(exampleFilename(13))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay13Part1(t *testing.T) {
	const want uint = 5588
	lines, err := linesFromFilename(filename(13))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay13Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(13))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		_ = Day13(lines, true)
	}
}
