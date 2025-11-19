package adventofcode2022

import (
	"fmt"
	"testing"
)

var containsTests = []struct {
	a1, a2, b1, b2 int
	contained      bool
	overlap        bool
}{
	{1, 4, 2, 3, true, true},
	{1, 1, 1, 1, true, true},
	{1, 1, 1, 3, true, true},
	{1, 3, 1, 1, true, true},
	{1, 1, 2, 2, false, false},
	{2, 2, 1, 1, false, false},
	{0, 0, 1, 9, false, false},
	{1, 9, 0, 0, false, false},
	{5, 7, 7, 9, false, true},
	{2, 8, 3, 7, true, true},
	{6, 6, 4, 6, true, true},
	{2, 6, 4, 8, false, true},
}

func TestContains(t *testing.T) {
	for _, tt := range containsTests {
		id := fmt.Sprintf("[%d-%d] <-> [%d-%d]", tt.a1, tt.a2,
			tt.b1, tt.b2)
		t.Run(id, func(t *testing.T) {
			got := Contains(tt.a1, tt.a2, tt.b1, tt.b2)
			if tt.contained != got {
				t.Fatalf("want %t but got %t", tt.contained, got)
			}
		})
	}
}

func TestDay04Part1Example(t *testing.T) {
	testLinesErr(t, 4, exampleFilename, true, Day04, 2)
}

func TestDay04Part1(t *testing.T) {
	testLinesErr(t, 4, filename, true, Day04, 580)
}

func TestDay04Part2Example(t *testing.T) {
	testLinesErr(t, 4, exampleFilename, false, Day04, 4)
}

func TestDay04Part2(t *testing.T) {
	testLinesErr(t, 4, filename, false, Day04, 895)
}

func BenchmarkDay04Part1(b *testing.B) {
	benchLinesErr(b, 4, true, Day04)
}

func BenchmarkDay04Part2(b *testing.B) {
	benchLinesErr(b, 4, false, Day04)
}
