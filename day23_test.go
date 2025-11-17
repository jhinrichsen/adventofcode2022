package adventofcode2022

import "testing"

// Mini example - Part 2
func TestDay23Part2Example0(t *testing.T) {
	const want uint = 4
	got := Day23([]string{
		".....",
		"..##.",
		"..#..",
		".....",
		"..##.",
		".....",
	}, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1Example(t *testing.T) {
	const want uint = 110
	lines, err := linesFromFilename(exampleFilename(23))
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part1(t *testing.T) {
	const want uint = 3689
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part2Example(t *testing.T) {
	const want uint = 20
	lines, err := linesFromFilename(exampleFilename(23))
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part2(t *testing.T) {
	const want uint = 965
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		t.Fatal(err)
	}
	got := Day23(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day23(lines, true)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	lines, err := linesFromFilename(filename(23))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day23(lines, false)
	}
}
