package adventofcode2021

import "testing"

func TestDay02Part1Example(t *testing.T) {
	const want = 15
	lines, err := linesFromFilename(exampleFilename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part1(t *testing.T) {
	const want = 12794
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day02(lines)
	}
}
