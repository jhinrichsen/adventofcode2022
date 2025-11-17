package adventofcode2022

import "testing"

func TestDay14Part1Example(t *testing.T) {
	const want uint = 24
	lines, err := linesFromFilename(exampleFilename(14))
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part1(t *testing.T) {
	const want uint = 779
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part2Example(t *testing.T) {
	const want uint = 93
	lines, err := linesFromFilename(exampleFilename(14))
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Part2(t *testing.T) {
	const want uint = 27426
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		t.Fatal(err)
	}
	got := Day14(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay14Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		b.Skip("input file not found")
	}
	b.ResetTimer()
	for range b.N {
		_ = Day14(lines, true)
	}
}

func BenchmarkDay14Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(14))
	if err != nil {
		b.Skip("input file not found")
	}
	b.ResetTimer()
	for range b.N {
		_ = Day14(lines, false)
	}
}
