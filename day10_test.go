package adventofcode2022

import "testing"

func TestDay10Example(t *testing.T) {
	const want = 13140
	lines, err := linesFromFilename(exampleFilename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10(t *testing.T) {
	const want = 15140
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		t.Fatal(err)
	}
	got := Day10(lines)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10(b *testing.B) {
	lines, err := linesFromFilename(filename(10))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day10(lines)
	}
}
