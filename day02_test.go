package adventofcode2022

import "testing"

func TestDay02Part1Example(t *testing.T) {
	const want = 15
	lines, err := linesFromFilename(exampleFilename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part1(t *testing.T) {
	const want = 13052
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines, true)
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
		_ = Day02(lines, true)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day02(lines, false)
	}
}

func TestDay02Part2Example(t *testing.T) {
	const want = 12
	lines, err := linesFromFilename(exampleFilename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = 13693
	lines, err := linesFromFilename(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day02(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
