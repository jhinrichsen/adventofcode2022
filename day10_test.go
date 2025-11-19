package adventofcode2022

import (
	"strings"
	"testing"

	"gitlab.com/jhinrichsen/aococr"
)

func TestDay10Part1Example(t *testing.T) {
	const want uint = 13140
	lines := linesFromFilename(t, exampleFilename(10))
	got, _ := Day10(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want uint = 15140
	lines := linesFromFilename(t, filename(10))
	got, _ := Day10(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(10))
	for range b.N {
		_, _ = Day10(lines, true)
	}
}

// TestDay10Part2 parses the ASCII CRT output into a string using aococr
// and verifies the expected answer for Part 2.
func TestDay10Part2(t *testing.T) {
	const want = "BPJAZGAP"

	lines := linesFromFilename(t, filename(10))
	_, crt := Day10(lines, false)
	got, err := aococr.ParseLetters(strings.Join(crt, "\n"), map[rune]bool{'#': true})
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(10))
	for range b.N {
		_, _ = Day10(lines, false)
	}
}
