package adventofcode2022

import (
	"fmt"
	"testing"
)

func TestDay24Part1Example(t *testing.T) {
	const want uint = 18
	lines := linesFromFilename(t, exampleFilename(24))
	puzzle := NewDay24(lines)
	got := Day24(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part1(t *testing.T) {
	const want uint = 299
	lines := linesFromFilename(t, filename(24))
	puzzle := NewDay24(lines)
	got := Day24(puzzle, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part2Example(t *testing.T) {
	const want uint = 54
	lines := linesFromFilename(t, exampleFilename(24))
	puzzle := NewDay24(lines)
	got := Day24(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part2(t *testing.T) {
	const want uint = 899
	lines := linesFromFilename(t, filename(24))
	puzzle := NewDay24(lines)
	got := Day24(puzzle, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// TestDay24Part1AllMinutes verifies blizzard configuration at all minutes 0-18
func TestDay24Part1AllMinutes(t *testing.T) {
	lines := linesFromFilename(t, exampleFilename(24))
	puzzle := NewDay24(lines)

	tests := []struct {
		minute   uint
		expected []string
	}{
		{0, []string{"#.######", "#>>.<^<#", "#.<..<<#", "#>v.><>#", "#<^v^^>#", "######.#"}},
		{1, []string{"#.######", "#.>3.<.#", "#<..<<.#", "#>2.22.#", "#>v..^<#", "######.#"}},
		{2, []string{"#.######", "#.2>2..#", "#.^22^<#", "#.>2.^>#", "#.>..<.#", "######.#"}},
		{3, []string{"#.######", "#<^<22.#", "#.2<.2.#", "#><2>..#", "#..><..#", "######.#"}},
		{4, []string{"#.######", "#.<..22#", "#<<.<..#", "#<2.>>.#", "#.^22^.#", "######.#"}},
		{5, []string{"#.######", "#2.v.<>#", "#<.<..<#", "#.^>^22#", "#.2..2.#", "######.#"}},
		{6, []string{"#.######", "#>2.<.<#", "#.2v^2<#", "#>..>2>#", "#<....>#", "######.#"}},
		{7, []string{"#.######", "#.22^2.#", "#<v.<2.#", "#>>v<>.#", "#>....<#", "######.#"}},
		{8, []string{"#.######", "#.<>2^.#", "#..<<.<#", "#.22..>#", "#.2v^2.#", "######.#"}},
		{9, []string{"#.######", "#<.2>>.#", "#.<<.<.#", "#>2>2^.#", "#.v><^.#", "######.#"}},
		{10, []string{"#.######", "#.2..>2#", "#<2v2^.#", "#<>.>2.#", "#..<>..#", "######.#"}},
		{11, []string{"#.######", "#2^.^2>#", "#<v<.^<#", "#..2.>2#", "#.<..>.#", "######.#"}},
		{12, []string{"#.######", "#>>.<^<#", "#.<..<<#", "#>v.><>#", "#<^v^^>#", "######.#"}},
		{13, []string{"#.######", "#.>3.<.#", "#<..<<.#", "#>2.22.#", "#>v..^<#", "######.#"}},
		{14, []string{"#.######", "#.2>2..#", "#.^22^<#", "#.>2.^>#", "#.>..<.#", "######.#"}},
		{15, []string{"#.######", "#<^<22.#", "#.2<.2.#", "#><2>..#", "#..><..#", "######.#"}},
		{16, []string{"#.######", "#.<..22#", "#<<.<..#", "#<2.>>.#", "#.^22^.#", "######.#"}},
		{17, []string{"#.######", "#2.v.<>#", "#<.<..<#", "#.^>^22#", "#.2..2.#", "######.#"}},
		{18, []string{"#.######", "#>2.<.<#", "#.2v^2<#", "#>..>2>#", "#<....>#", "######.#"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Minute%d", tt.minute), func(t *testing.T) {
			actual := puzzle.renderGrid(tt.minute)
			for y := 0; y < len(tt.expected); y++ {
				if actual[y] != tt.expected[y] {
					t.Errorf("Minute %d, row %d:\n  want: %s\n  got:  %s", tt.minute, y, tt.expected[y], actual[y])
				}
			}
		})
	}
}

func BenchmarkDay24Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(24))
	puzzle := NewDay24(lines)
	b.ResetTimer()
	for range b.N {
		_ = Day24(puzzle, true)
	}
}

func BenchmarkDay24Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(24))
	puzzle := NewDay24(lines)
	b.ResetTimer()
	for range b.N {
		_ = Day24(puzzle, false)
	}
}
